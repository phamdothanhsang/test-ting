import { table } from "console";
import { Kols } from '../types';

interface KolCardProps {
    kol: Kols;
}
function KolCard({ kol }: KolCardProps) {
    return ( 
       
    <tr>
       <td className="tbody-border">{kol.kolID}</td>
            <td className="tbody-border">{kol.code}</td>
            <td className="tbody-border">{kol.language}</td>
            <td className="tbody-border">{kol.education}</td>
            <td className="tbody-border">{kol.expectedSalary}</td>
            <td className="tbody-border">{kol.expectedSalaryEnable ? "Yes" : "No"}</td>
            <td className="tbody-border">{kol.channelSettingTypeID}</td>
            <td className="tbody-border"><img src={kol.iDFrontURL} alt="ID Front" width="50" /></td>
            <td className="tbody-border"><img src={kol.iDBackURL} alt="ID Back" width="50" /></td>
            <td className="tbody-border"><img src={kol.portraitURL} alt="Portrait" width="50" /></td>
            <td className="tbody-border">{kol.rewardID}</td>
            <td className="tbody-border">{kol.paymentMethodID}</td>
            <td className="tbody-border">{kol.testimonialsID}</td>
            <td className="tbody-border">{kol.verificationStatus ? "Verified" : "Not Verified"}</td>
            <td className="tbody-border">{kol.enabled ? "Enabled" : "Disabled"}</td>
            <td className="tbody-border">{kol.activeDate}</td>
            <td className="tbody-border">{kol.active ? "Active" : "Inactive"}</td>
            <td className="tbody-border">{kol.createdBy}</td>
            <td className="tbody-border">{kol.createdDate}</td>
            <td className="tbody-border">{kol.modifiedBy}</td>
            <td className="tbody-border">{kol.modifiedDate}</td>
            <td className="tbody-border">{kol.isRemove ? "Removed" : "Not Removed"}</td>
            <td className="tbody-border">{kol.isOnBoarding ? "Onboarding" : "Not Onboarding"}</td>
            <td className="tbody-border">{kol.portraitRightURL}</td>
            <td className="tbody-border">{kol.portraitLeftURL}</td>
            <td className="tbody-border">{kol.livenessStatus ? "Live" : "Not Live"}</td>
    </tr>

     );
}
 
export default KolCard;