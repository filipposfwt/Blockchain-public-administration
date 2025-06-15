import { $api_dms_diplomas } from "@/composables/api";

import type {
    Diploma,
    DiplomaRequest,
    DiplomaVerificationResponse
  } from "@/repository/modules/diplomas/types";
  

  const getDiplomas = async (): Promise<Diploma[]> => {
    const response = await $api_dms_diplomas.get<Diploma[]>(
      `/diplomas`
    );
    return response.data;
  };  

  const getDiploma = async (hash: string): Promise<Diploma> => {
    const response = await $api_dms_diplomas.get<Diploma>(
      `/diplomas/${hash}`
    );
    return response.data;
  };

  const createDiploma = async (diploma: DiplomaRequest): Promise<Diploma> => {
    const response = await $api_dms_diplomas.post<Diploma>(
      `/diplomas`,
      diploma 
    );
    return response.data;
  };

  const verifyDiploma = async (hash: string): Promise<DiplomaVerificationResponse> => {
    console.log('Making API call to verify diploma with hash:', hash);
    const response = await $api_dms_diplomas.get<DiplomaVerificationResponse>(
      `/diplomas/${hash}`
    );
    console.log('API response:', response);
    return response.data;
  };



  
  export { createDiploma, verifyDiploma, getDiplomas, getDiploma };
  