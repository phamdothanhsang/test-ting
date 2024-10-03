// services/KolService.ts
import axios from "axios";

interface KolRequest {
  pageIndex?: number | null;
  pageSize?: number | null;
}

export const fetchKols = async (params: KolRequest) => {
  try {
    const response = await axios.get("http://localhost:8081/kols", {
      params,
    });
    return response.data;
  } catch (error) {
    console.error("Error fetching KOL data:", error);
    throw error;
  }
};
