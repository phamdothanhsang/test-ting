import React from "react";
import { Kols } from "@/models/kol";

interface ModalProps {
  isOpen: boolean;
  onClose: () => void;
  kol: Kols; // Use the Kols interface here
}

const Modal: React.FC<ModalProps> = ({ isOpen, onClose, kol }) => {
  if (!isOpen) return null;

  return (
    <div className="modal-overlay" onClick={onClose}>
      <div className="modal-content" onClick={(e) => e.stopPropagation()}>
        <button className="close-button" onClick={onClose}>
          X
        </button>
        <div className="modal-body">
          <div className="modal-images">
            <img src={kol.PortraitURL} alt="Profile" />
            <img src={kol.IDFrontURL} alt="ID Front" />
            <img src={kol.IDBackURL} alt="ID Back" />
          </div>
          <h2>KOL CODE: {kol.Code}</h2>
          <p>Language: {kol.Language}</p>
          <p>Education: {kol.Education}</p>
          <p>Expected Salary: ${kol.ExpectedSalary}</p>
          <p>Payment Method: {kol.PaymentMethodID}</p>
          <p>Verification Status: {kol.VerificationStatus} ✅</p>
          <p>Liveness Status: {kol.LivenessStatus} ✅</p>
          <p>
            Onboarding Status:{" "}
            {kol.IsOnBoarding ? "Completed ✅" : "Not Completed"}
          </p>
          <p>Created: {kol.CreatedDate.toLocaleDateString()}</p>
          <p>Last Modified: {kol.ModifiedDate.toLocaleDateString()}</p>
        </div>
      </div>
    </div>
  );
};

export default Modal;
