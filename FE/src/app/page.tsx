"use client"

import React, { useEffect, useState } from 'react';

// Import CSS file, you can use CSS module if you want
import './page.css'

interface Kols {
    KolID: number;
    UserProfileID: number;
    Language: string;
    Education: string;
    ExpectedSalary: number;
    ExpectedSalaryEnable: boolean;
    ChannelSettingTypeID: number;
    IDFrontURL: string;
    IDBackURL: string;
    PortraitURL: string;
    RewardID: number;
    PaymentMethodID: number;
    TestimonialsID: number;
    VerificationStatus: boolean;
    Enabled: boolean;
    ActiveDate: Date;
    Active: boolean;
    CreatedBy: string;
    CreatedDate: Date;
    ModifiedBy: string;
    ModifiedDate: Date;
    IsRemove: boolean;
    IsOnBoarding: boolean;
    Code: string;
    PortraitRightURL: string;
    PortraitLeftURL: string;
    LivenessStatus: boolean;
}

const Page = () => {
    const [Kols, setKols] = useState<Kols[]>([]);

    useEffect(() => {
        const fetchData = async () => {
            try {
                const response = await fetch('http://localhost:8081/kols'); 
                const data = await response.json();
					 const KOL = data.kol;
                // Chuyển đổi dữ liệu từ API về kiểu Kols
                const transformedData: Kols[] = KOL.map((item: any) => ({
                    KolID: item.kolID,
                    UserProfileID: item.userProfileID,
                    Language: item.language,
                    Education: item.education,
                    ExpectedSalary: item.expectedSalary,
                    ExpectedSalaryEnable: item.expectedSalaryEnable,
                    ChannelSettingTypeID: item.channelSettingTypeID,
                    IDFrontURL: item.iDFrontURL,
                    IDBackURL: item.iDBackURL,
                    PortraitURL: item.portraitURL,
                    RewardID: item.rewardID,
                    PaymentMethodID: item.paymentMethodID,
                    TestimonialsID: item.testimonialsID,
                    VerificationStatus: item.verificationStatus,
                    Enabled: item.enabled,
                    ActiveDate: new Date(item.activeDate),
                    Active: item.active,
                    CreatedBy: item.createdBy,
                    CreatedDate: new Date(item.createdDate),
                    ModifiedBy: item.modifiedBy,
                    ModifiedDate: new Date(item.modifiedDate),
                    IsRemove: item.isRemove,
                    IsOnBoarding: item.isOnBoarding,
                    Code: item.code,
                    PortraitRightURL: item.portraitRightURL,
                    PortraitLeftURL: item.portraitLeftURL,
                    LivenessStatus: item.livenessStatus
                }));

                setKols(transformedData);
            } catch (error) {
                console.error("Error fetching data:", error);
            }
        };

        fetchData();
    }, []);

    return (
		<div>
				<div className='header bg-warning p-2'>
					<h1 className='header text-center'>List KOL</h1>
				</div>
          
            <div className='listKOL mt-5 position-relative'>
                <table className="table table-striped table-warning">
                    <thead>
                        <tr>
                            <th>KOL ID</th>
                            <th>UserProfile ID</th>
                            <th>Language</th>
                            <th>Education</th>
                            <th>Expected Salary</th>
                            <th>Expected Salary Enable</th>
                            <th>Channel Setting Type ID</th>
                            <th>ID Front URL</th>
                            <th>ID Back URL</th>
                            <th>Portrait URL</th>
                            <th>Reward ID</th>
                            <th>Payment Method ID</th>
                            <th>Testimonials ID</th>
                            <th>Verification Status</th>
                            <th>Enabled</th>
                            <th>Active Date</th>
                            <th>Active</th>
                            <th>Created By</th>
                            <th>Created Date</th>
                            <th>Modified By</th>
                            <th>Modified Date</th>
                            <th>Is Remove</th>
                            <th>Is On Boarding</th>
                            <th>Code</th>
                            <th>Portrait Right URL</th>
                            <th>Portrait Left URL</th>
                            <th>Liveness Status</th>
                        </tr>
                    </thead>
                    <tbody>
                        {Kols.map(kol => (
                            <tr key={kol.KolID}>
                                <td>{kol.KolID}</td>
                                <td>{kol.UserProfileID}</td>
                                <td>{kol.Language}</td>
                                <td>{kol.Education}</td>
                                <td>{kol.ExpectedSalary}</td>
                                <td>{kol.ExpectedSalaryEnable ? "Yes" : "No"}</td>
                                <td>{kol.ChannelSettingTypeID}</td>
                                <td><a href={kol.IDFrontURL} >View</a></td>
                                <td><a href={kol.IDBackURL} >View</a></td>
                                <td><a href={kol.PortraitURL} >View</a></td>
                                <td>{kol.RewardID}</td>
                                <td>{kol.PaymentMethodID}</td>
                                <td>{kol.TestimonialsID}</td>
                                <td>{kol.VerificationStatus ? "Verified" : "Not Verified"}</td>
                                <td>{kol.Enabled ? "Yes" : "No"}</td>
                                <td>{kol.ActiveDate.toLocaleDateString()}</td>
                                <td>{kol.Active ? "Yes" : "No"}</td>
                                <td>{kol.CreatedBy}</td>
                                <td>{kol.CreatedDate.toLocaleDateString()}</td>
                                <td>{kol.ModifiedBy}</td>
                                <td>{kol.ModifiedDate.toLocaleDateString()}</td>
                                <td>{kol.IsRemove ? "Yes" : "No"}</td>
                                <td>{kol.IsOnBoarding ? "Yes" : "No"}</td>
                                <td>{kol.Code}</td>
                                <td><a href={kol.PortraitRightURL}>View</a></td>
                                <td><a href={kol.PortraitLeftURL}>View</a></td>
                                <td>{kol.LivenessStatus ? "Live" : "Not Live"}</td>
                            </tr>
                        ))}
                    </tbody>
                </table>
            </div>
        </div>
    );
};

export default Page;
