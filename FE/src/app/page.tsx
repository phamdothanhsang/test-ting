"use client";

import React, { useEffect, useState } from "react";

// * Import CSS file, you can use CSS module if you want
// ! Change your CSS inside this file
import "./page.css";
import { Kols, transformKol } from "@/models/kol";
import { fetchKols } from "@/services/kols-services";
import KolsCard from "@/components/kolsCard";
import Modal from "@/components/Modal";
import Loading from "@/components/loading";
import TableHead from "@/components/tableHead";

const Page = () => {
  // * Use useState to store Kols from API
  // ! (if you have more optimized way to store data, PLEASE FEELS FREE TO CHANGE)
  const [Kols, setKols] = useState<Kols[]>([]);
  const [loading, setLoading] = useState(true);

  const [pageIndex, setPageIndex] = useState(1);
  const [selectedKol, setSelectedKol] = useState<Kols | null>(null);
  const [isModalOpen, setIsModalOpen] = useState(false);

  const handleRowClick = (kol: Kols) => {
    setSelectedKol(kol);
    setIsModalOpen(true);
  };

  const handleCloseModal = () => {
    setIsModalOpen(false);
    setSelectedKol(null);
  };
  const getKols = async () => {
    setLoading(true);
    try {
      const data = await fetchKols({ pageIndex: pageIndex, pageSize: 10 }); // Adjust as needed
      const transformedData: Kols[] = data.kol.map((item: any) =>
        transformKol(item)
      );
      setKols(transformedData || []); // Adjust based on your response
    } catch (error) {
      console.error("Error fetching KOLs:", error);
    } finally {
      setLoading(false);
    }
  };

  const handleNextPage = () => {
    setPageIndex(pageIndex + 1);
  };

  useEffect(() => {
    getKols();
  }, [pageIndex]);

  return (
    <>
      <div className="table-container">
        <h1>KOLs Information</h1>
        {loading ? (
          <div>
            <table className="kol-table">
              <TableHead />
            </table>
            <Loading />
          </div>
        ) : (
          <table className="kol-table">
            <TableHead />
            <tbody>
              {Kols.map((kol) => (
                <KolsCard
                  onClick={() => handleRowClick(kol)}
                  key={kol.KolID}
                  kolID={kol.KolID}
                  code={kol.Code}
                  language={kol.Language}
                  expectedSalary={kol.ExpectedSalary}
                  verificationStatus={kol.VerificationStatus}
                />
              ))}
            </tbody>
          </table>
        )}
        <div className="next-page-div">
          <button onClick={handleNextPage} className="green-button">
            Next Page
          </button>
        </div>
        {selectedKol && (
          <Modal
            isOpen={isModalOpen}
            onClose={handleCloseModal}
            kol={selectedKol}
          />
        )}
      </div>
    </>
  );
};

export default Page;
