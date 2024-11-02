"use client"

import React from 'react';

// * Import CSS file, you can use CSS module if you want
// ! Change your CSS inside this file
import './page.css'

import KolsTable from './components/kolsTable';

const Page = () => {
    return (
        <div>	
            <h1 className='header'>Kol Table</h1>
			<KolsTable/>
        </div>
    )
};

export default Page;