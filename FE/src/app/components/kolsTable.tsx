import React, {useEffect, useState} from "react";
import {Kols} from "../lib/definitions"


export default function KolsTable() {
    // * Use useState to store Kols from API 
    // ! (if you have more optimized way to store data, PLEASE FEELS FREE TO CHANGE)
	const [Kols , setKols] = useState<Kols[]>([]);  

    // * Fetch API over here 
    // * Use useEffect to fetch data from API 
    useEffect(() => {
		fetch('http://localhost:8081/kols?pageIndex=1&pageSize=100', {
			'method': 'GET',
			headers: {
				'Content-Type': 'application/json',

			},
		})
		.then(resp => resp.json())
		.then(resp => setKols(resp["KolInformation"]))
		.catch(error => console.log(error))
    }, []);

    return (
        <div className="overflow-auto hover:overflow-scroll kolTable max-h-full">
            <table className=" table-auto overflow-scroll w-full">
            <thead>
                <tr className='bg-gray-100 text-left'>
                    <th scope="col" className="px-4 py-2">Kol ID</th>
                    <th scope="col" className="px-4 py-2">User Profile ID</th>
                    <th scope="col" className="px-4 py-2">Language</th>
                    <th scope="col" className="px-4 py-2">Education</th>
                    <th scope="col" className="px-4 py-2">Expected Salary</th>
                    <th scope="col" className="px-4 py-2">Expected Salary Enable</th>
                    <th scope="col" className="px-4 py-2">Channel Setting Type ID</th>
                    <th scope="col" className="px-4 py-2">ID Front URL</th>
                    <th scope="col" className="px-4 py-2">ID Back URL</th>
                    <th scope="col" className="px-4 py-2">Portrait URL</th>
                    <th scope="col" className="px-4 py-2">Reward ID</th>
                    <th scope="col" className="px-4 py-2">Payment Method ID</th>
                    <th scope="col" className="px-4 py-2">Testimonials ID</th>
                    <th scope="col" className="px-4 py-2">Verification Status</th>
                    <th scope="col" className="px-4 py-2">Enabled</th>
                    <th scope="col" className="px-4 py-2">Active Date</th>
                    <th scope="col" className="px-4 py-2">Active</th>
                    <th scope="col" className="px-4 py-2">Created By</th>
                    <th scope="col" className="px-4 py-2">Created Date</th>
                    <th scope="col" className="px-4 py-2">Modified By</th>
                    <th scope="col" className="px-4 py-2">Modified Date</th>
                    <th scope="col" className="px-4 py-2">Is Remove</th>
                    <th scope="col" className="px-4 py-2">Is On Boarding</th>
                    <th scope="col" className="px-4 py-2">Code</th>
                    <th scope="col" className="px-4 py-2">Portrait Right URL</th>
                    <th scope="col" className="px-4 py-2">Portrait Left URL</th>
                    <th scope="col" className="px-4 py-2">Liveness Status</th>
                </tr>
            </thead>
            <tbody>
            {Kols && Kols.map(Kol => {
                Kol.ActiveDate = new Date(Date.parse(Kol.ActiveDate.toString()))
                Kol.CreatedDate = new Date(Date.parse(Kol.CreatedDate.toString()))
                Kol.ModifiedDate = new Date(Date.parse(Kol.ModifiedDate.toString()))
                return (
                    <tr className="border-b" key={Kol.KolID}>
                        <td className="px-4 py-2">{Kol.KolID}</td>
                        <td className="px-4 py-2">{Kol.UserProfileID}</td>
                        <td className="px-4 py-2">{Kol.Language}</td>
                        <td className="px-4 py-2">{Kol.Education}</td>
                        <td className="px-4 py-2">{Kol.ExpectedSalary}</td>
                        <td className="px-4 py-2">{Kol.ExpectedSalaryEnable.toString()}</td>
                        <td className="px-4 py-2">{Kol.ChannelSettingTypeID}</td>
                        <td className="px-4 py-2">{Kol.IDFrontURL}</td>
                        <td className="px-4 py-2">{Kol.IDBackURL}</td>
                        <td className="px-4 py-2">{Kol.PortraitURL}</td>
                        <td className="px-4 py-2">{Kol.RewardID}</td>
                        <td className="px-4 py-2">{Kol.PaymentMethodID}</td>
                        <td className="px-4 py-2">{Kol.TestimonialsID}</td>
                        <td className="px-4 py-2">{Kol.VerificationStatus}</td>
                        <td className="px-4 py-2">{Kol.Enabled.toString()}</td>
                        <td className="px-4 py-2">{Kol.ActiveDate.toISOString()}</td>
                        <td className="px-4 py-2">{Kol.Active.toString()}</td>
                        <td className="px-4 py-2">{Kol.CreatedBy}</td>
                        <td className="px-4 py-2">{Kol.CreatedDate.toISOString()}</td>
                        <td className="px-4 py-2">{Kol.ModifiedBy}</td>
                        <td className="px-4 py-2">{Kol.ModifiedDate.toISOString()}</td>
                        <td className="px-4 py-2">{Kol.IsRemove.toString()}</td>
                        <td className="px-4 py-2">{Kol.IsOnBoarding.toString()}</td>
                        <td className="px-4 py-2">{Kol.Code}</td>
                        <td className="px-4 py-2">{Kol.PortraitRightURL}</td>
                        <td className="px-4 py-2">{Kol.PortraitLeftURL}</td>
                        <td className="px-4 py-2">{Kol.LivenessStatus}</td>
                     </tr>

                )
            })}
            </tbody>
            </table>
        </div>
    )
}