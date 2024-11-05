import KolCard from "./kolCard";
import { Kols } from "../types";

interface ListOfKols {
    kols: Kols[];
}
const KolList = ({kols}: ListOfKols) => {
    return ( 
        <table>
            <thead>
                <tr>
                <th className="thead-border">Kol ID</th>
                    <th className="thead-border">Code</th>
                    <th className="thead-border">Language</th>
                    <th className="thead-border">Education</th>
                    <th className="thead-border">Expected Salary</th>
                    <th className="thead-border">Expected Salary Enable</th>
                    <th className="thead-border">Channel Setting Type ID</th>
                    <th className="thead-border">ID Front URL</th>
                    <th className="thead-border">ID Back URL</th>
                    <th className="thead-border">Portrait URL</th>
                    <th className="thead-border">Reward ID</th>
                    <th className="thead-border">Payment Method ID</th>
                    <th className="thead-border">Testimonials ID</th>
                    <th className="thead-border">Verification Status</th>
                    <th className="thead-border">Enabled</th>
                    <th className="thead-border">Active Date</th>
                    <th className="thead-border">Active</th>
                    <th className="thead-border">Created By</th>
                    <th className="thead-border">Created Date</th>
                    <th className="thead-border">Modified By</th>
                    <th className="thead-border">Modified Date</th>
                    <th className="thead-border">Is Removed</th>
                    <th className="thead-border">Is Onboarding</th>
                    <th className="thead-border">Portrait Right URL</th>
                    <th className="thead-border">Portrait Left URL</th>
                    <th className="thead-border">Liveness Status</th>
                </tr>
            </thead>
            <tbody>
                {kols.map((kol) => (
                    <KolCard key={kol.kolID} kol={kol} />
                ))}
            </tbody>
        </table>
     );
}
 
export default KolList;