import React, { useEffect, useState } from "react";

type TableHeadType = {};

const TableHead: React.FC<TableHeadType> = ({}) => {
  return (
    <thead>
      <tr>
        <th>KOL ID</th>
        <th>KOL CODE</th>
        <th>Language</th>
        <th>Expected Salary</th>
        <th>Verification Status</th>
      </tr>
    </thead>
  );
};

export default TableHead;
