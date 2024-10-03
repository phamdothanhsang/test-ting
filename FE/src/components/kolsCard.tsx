import React, { useEffect, useState } from "react";

type KolsCardType = {
  kolID: number;
  code: string;
  language: string;
  expectedSalary: number;
  verificationStatus: string;
  onClick: () => void;
};

const KolsCard: React.FC<KolsCardType> = ({
  kolID,
  code,
  language,
  expectedSalary,
  verificationStatus,
  onClick,
}) => {
  return (
    <tr className="card-clickable" onClick={onClick}>
      <td>{kolID}</td>
      <td>{code}</td>
      <td>{language}</td>
      <td>{expectedSalary}</td>
      <td className="verified">{verificationStatus}</td>
    </tr>
  );
};

export default KolsCard;
