import React, { useEffect, useState } from "react";

type LoadingType = {};

const Loading: React.FC<LoadingType> = ({}) => {
  return (
    <div className="loader-container">
      <div className="loader"></div>
    </div>
  );
};

export default Loading;
