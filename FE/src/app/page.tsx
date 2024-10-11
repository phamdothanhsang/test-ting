"use client"

import React, {useEffect, useState} from 'react';

import {Table, TableHeader, TableColumn, TableBody, TableRow, TableCell, getKeyValue} from "@nextui-org/react";

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
	}
];

const Page = () => {
    // * Use useState to store Kols from API 
    // ! (if you have more optimized way to store data, PLEASE FEELS FREE TO CHANGE)
	const [KolsList , setKols] = useState<Kols[]>([]);  

    // * Fetch API over here 
    // * Use useEffect to fetch data from API 
	
    useEffect(() => {
		// Could be GET or POST/PUT/PATCH/DELETE
		fetch('http://localhost:8081/kols?page_size=5&page_index=1')
			.then(res => res.json())
			.then((res) => {
				setKols(res["kol"])
				// console.log(KolsList)
			})
	}, [])
	


/* { status: 'ok', method: 'GET' } */

	
	return (
		<Table aria-label="Example table with dynamic content">
			<TableHeader columns={columns}>
				{(column) => <TableColumn key={column.key}>{column.label}</TableColumn>}
			</TableHeader>
			<TableBody items={KolsList}>
				{(item) => (
				<TableRow key={item.kolID}>
					{(columnKey) => <TableCell>{getKeyValue(item, columnKey)}</TableCell>}
				</TableRow>
				)}
			</TableBody>
		</Table>
	);
}

export default Page;