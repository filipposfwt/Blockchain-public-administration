export type Diploma = {
  id: string;
  type: string;
  issuer: string;
  holder: string;
  issuedDate: string;
  expiryDate: string;
  degreeName: string;
  university: string;
  gradePoint: number;
}

export type DiplomaRequest = {
  diplomaData: Diploma;
  issuerDID: string;
  privateKey: string;
}

export type DiplomaVerificationResponse = {
  verified: boolean;
  issuerDID: string;
}


  

  
  
  