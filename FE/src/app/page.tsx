'use client';

import React, { useEffect, useState } from 'react';
import axiosInstance from './api';
import './page.css';

interface Kol {
  kolID: number;
  userProfileID: number;
  language: string;
  education: string;
  expectedSalary: number;
  expectedSalaryEnable: boolean;
  channelSettingTypeID: number;
  iDFrontURL: string;
  iDBackURL: string;
  portraitURL: string;
  rewardID: number;
  paymentMethodID: number;
  testimonialsID: number;
  verificationStatus: string;
  enabled: boolean;
  activeDate: string;
  createdBy: string;
  createdDate: string;
  modifiedBy: string;
  modifiedDate: string;
  isRemove: boolean;
  isOnBoarding: boolean;
  code: string;
  portraitRightURL: string;
  portraitLeftURL: string;
  livenessStatus: string;
}

interface KolViewModel {
  KolInformation: Kol[];
  totalCount: number;
  totalPages: number;
}

const Page = () => {
  const [kols, setKols] = useState<Kol[]>([]);
  const [page, setPage] = useState(1);
  const [pageSize] = useState(10);
  const [totalCount, setTotalCount] = useState(0);
  const [totalPages, setTotalPages] = useState(0);
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    const fetchKols = async () => {
      setLoading(true);
      try {
        const response = await axiosInstance.get<KolViewModel>('/kols', {
          params: {
            page: page,
            pageSize: pageSize,
          },
        });
        setKols(response.data.KolInformation);
        setTotalCount(response.data.totalCount);
        setTotalPages(response.data.totalPages);
      } catch (error) {
        console.error('Error fetching KOL data:', error);
      } finally {
        setLoading(false);
      }
    };
    fetchKols();
  }, [page, pageSize]);

  const handlePageChange = (newPage: number) => {
    if (newPage >= 1 && newPage <= totalPages) {
      setPage(newPage);
      window.scrollTo(0, 0);
    }
  };

  return (
    <div className="container">
      <h1>KOL List</h1>
      {loading ? (
        <p>Loading...</p>
      ) : kols.length > 0 ? (
        kols.map((kol) => (
          <div key={kol.kolID || kol.code} className="kol-item">
            <h2>KOL Information</h2>
            <p><strong>KOL ID:</strong> {kol.kolID}</p>
            <p><strong>User Profile ID:</strong> {kol.userProfileID}</p>
            <p><strong>Language:</strong> {kol.language}</p>
            <p><strong>Education:</strong> {kol.education}</p>
            <p><strong>Expected Salary:</strong> {kol.expectedSalaryEnable ? `$${kol.expectedSalary}` : 'Not displayed'}</p>
            <p><strong>Portrait URL:</strong> {kol.portraitURL}</p>
            <p><strong>Verification Status:</strong> {kol.verificationStatus}</p>
            <p><strong>Enabled:</strong> {kol.enabled ? 'Yes' : 'No'}</p>
            <p><strong>Active Date:</strong> {new Date(kol.activeDate).toLocaleString()}</p>
            <p><strong>Created By:</strong> {kol.createdBy}</p>
            <p><strong>Modified By:</strong> {kol.modifiedBy}</p>
            <p><strong>Liveness Status:</strong> {kol.livenessStatus ? 'Passed' : 'Failed'}</p>
          </div>
        ))
      ) : (
        <p>No KOLs found.</p>
      )}

      <div className="pagination">
        <button disabled={page === 1} onClick={() => handlePageChange(page - 1)}>
          Previous
        </button>
        <span>Page {page} of {totalPages}</span>
        <button disabled={page === totalPages} onClick={() => handlePageChange(page + 1)}>
          Next
        </button>
      </div>
    </div>
  );
};

export default Page;
