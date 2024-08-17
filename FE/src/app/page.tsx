"use client"

import React, {useEffect, useState} from 'react';

// * Import CSS file, you can use CSS module if you want
// ! Change your CSS inside this file
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
    // * Use useState to store Kols from API 
    // ! (if you have more optimized way to store data, PLEASE FEELS FREE TO CHANGE)
	const [Kols , setKols] = useState<Kols[]>([]);  

    // * Fetch API over here 
    // * Use useEffect to fetch data from API 
    useEffect(() => {

    }, []);

    return (
        <>
            <h1 className='header'>Implement component over here</h1>
        </>
    )
};

export default Page;