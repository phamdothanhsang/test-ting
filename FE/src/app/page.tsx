"use client"

import React, {useEffect, useState} from 'react';
import KolList from './components/kolList';
import { Kols } from './types';
// * Import CSS file, you can use CSS module if you want
// ! Change your CSS inside this file
import './page.css'
interface ApiResponse {
    result: string;
    errorMessage: string;
    pageIndex: number;
    pageSize: number;
    guid: string;
    totalCount: number;
    kol: Kols[]; // Mảng chứa các đối tượng KOL
}

const Page = () => {
    // * Use useState to store Kols from API 
    // ! (if you have more optimized way to store data, PLEASE FEELS FREE TO CHANGE)
	const [Kols , setKols] = useState<Kols[]>([]);  

    // * Fetch API over here 
    // * Use useEffect to fetch data from API 
    useEffect(() => {
		//Khi muốn phân trang hay gọi một số lượng data nhất định chỉ cần truyền biến pageIndex, pageSize vào đường dẫn
		//http://localhost:8081/kols?pageIndex=${pageIndex}&pageSize=${pageSize}
		//http://localhost:8081/kols?pageIndex=1&pageSize=15
		fetch('http://localhost:8081/kols')
            .then(response => response.json())
            .then((data: ApiResponse) => {
                if (data.result === 'Success') {
					console.log("Kols data:", data.kol);
                    setKols(data.kol); // Lưu mảng kols từ API vào state
                }
            })
            .catch(error => console.error("Error fetching data:", error));
    }, []);

    return (
        <>
            <h1 className='title font-bold text-xl'>WEALLNET'S TEST</h1>
			<main className=' bg-[#31511E] w-full py-8'>
				<div className='bg-[#F6FCDF] shadow-xl w-4/5 ml-auto mr-auto py-6 rounded-xl'>
					<h1 className='text-center text-2xl font-bold py-8'>List of Kols</h1>
					<div className="table-wrapper px-10 h-[500px]">
						<div className="table-scroll center">
							<KolList kols={Kols} />
						</div>
					</div>
					
				</div>
			</main>
			
			
        </>
    )
};

export default Page;