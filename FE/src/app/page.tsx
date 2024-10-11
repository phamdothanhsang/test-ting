"use client"

import React, {useEffect, useState} from 'react';

import {Table, TableHeader, TableColumn, TableBody, TableRow, TableCell, getKeyValue, Pagination} from "@nextui-org/react";

// * Import CSS file, you can use CSS module if you want
// ! Change your CSS inside this file
import './page.css'

interface Kols {
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
	verificationStatus: boolean;
	enabled: boolean;
	activeDate: Date;
	active: boolean;
	createdBy: string;
	createdDate: Date;
	modifiedBy: string;
	modifiedDate: Date;
	isRemove: boolean;
	isOnBoarding: boolean;
	code: string;
	portraitRightURL: string;
	portraitLeftURL: string;
	livenessStatus: boolean;
}

const columns = [
	{
	  key: "kolID",
	  label: "ID",
	},
	{
	  key: "userProfileID",
	  label: "Profile ID",
	},
	{
	  key: "language",
	  label: "Language",
	},
	{
		key: "education",
		label: "Education"
	},
	{
		key:"expectedSalary",
		label: "Expected Salary"
	},
	{
		key: "expectedSalaryEnable",
		label: "Expected Salary Enable"
	},
	{
		key:"channelSettingTypeID",
		label:"Channel Setting Type ID"
	},
	{
		key: "iDFrontURL",
		label: "FrontURL ID"
	},
	{
		key: "iDBackURL",
		label: "BackURL ID"
	},
	{
		key: "portraitURL",
		label: "Portrait URL"
	},
	{
		key: "rewardID",
		label: "Reward ID"
	},
	{
		key: "paymentMethodID",
		label: "Payment Method ID"
	},
	{
		key: "testimonialsID",
		label: "Testimonials ID"
	},
	{
		key: "verificationStatus",
		label: "Verification Status"
	},
	{
		key: "enabled",
		label: "Enabled"
	},
	{
		key: "activeDate",
		label: "Active Date"
	},
	{
		key: "active",
		label: "Active"
	},
	{
		key: "createdBy",
		label: "Created By"
	},
	{
		key: "createdDate",
		label:  "Created Date"
	},
	{
		key: "modifiedBy",
		label: "Modified By"
	},
	{
		key: "modifiedDate",
		label: "Modified Date"
	},
	{
		key: "isRemove",
		label: "Is Removed"
	},
	{
		key: "isOnBoarding",
		label: "Is Onboarding"
	},
	{
		key: "code",
		label: "Code"
	},
	{
		key:"portraitRightURL",
		label: "Portrait Right URL"
	},
	{
		key:"portraitLeftURL",
		label: "Portrait Left URL"
	},
	{
		key: "livenessStatus",
		label: "Liveness Status"
	}
];

const Page = () => {
    // * Use useState to store Kols from API 
    // ! (if you have more optimized way to store data, PLEASE FEELS FREE TO CHANGE)
	const [KolsList , setKols] = useState<Kols[]>([]);
	const pageSize = 10;
	const pageIndex = 1;

    // * Fetch API over here 
    // * Use useEffect to fetch data from API 
	
    useEffect(() => {
		// Could be GET or POST/PUT/PATCH/DELETE
		fetch(`http://localhost:8081/kols?page_size=${pageSize}&page_index=${pageIndex}`)
			.then(res => res.json())
			.then((res) => {
				setKols(res["kol"])
				// console.log(KolsList)
			})
	}, [])
	


/* { status: 'ok', method: 'GET' } */

	
	return (
		<>
		<h1 className="text-3xl font-bold underline text-center">
			KOLs Information
		</h1>
		<br></br>
		<Table 
			isStriped aria-label="test table" 
			classNames={{
        		base: "max-h-[300px]"
      		}}
		>
			<TableHeader columns={columns}>
				{(column) => <TableColumn key={column.key}>{column.label}</TableColumn>}
			</TableHeader>
			<TableBody items={KolsList}>
				{(item) => (
				<TableRow key={item.kolID}>
					{(columnKey) => <TableCell>{getKeyValue(item, columnKey).toString()}</TableCell>}
				</TableRow>
				)}
			</TableBody>
		</Table>
		</>
	);
}

export default Page;