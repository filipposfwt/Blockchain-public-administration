import { defineStore } from "pinia";
import type { Diploma } from "@/repository/modules/diplomas/types";
import {
  getDiplomas,
  createDiploma,
  verifyDiploma,
} from "@/repository/modules/diplomas/factory";

type Pagination = {
  current_page: string;
  page_size: string;
  total_count: string;
  total_pages: string;
};

const diplomasColumns = [
  {
    name: "type",
    required: true,
    label: "Type",
    align: "left",
    sortable: true,
  },    
  {
    name: "degreeName",
    label: "Degree Name",
    field: "degreeName",
    align: "center",
    sortable: true,
  },
  {
    name: "gradePoint",
    label: "Grade Point",
    field: "gradePoint",
    align: "center",
    sortable: true,
  },
  { name: "holder", label: "Holder", field: "holder", align: "center", sortable: true },
  {
    name: "hash",
    label: "Hash",
    field: "hash",
    align: "center",
    sortable: true,
  },
  {
    name: "issuedDate",
    label: "Issued Date",
    field: "issuedDate",
    align: "center",
    sortable: true,
  },


 

  {
    name: "actions",
    label: "",
    align: "center",
    sortable: false,
  },
];



export const useDiplomasStore = defineStore("diplomasStore", {
  state: () => {
    return {
      perPage: 20,
      diplomasColumns: diplomasColumns,
      verifyDiplomaDialog:false,
      addDiplomaDialog:false,
      viewDiplomaDialog:false,
      pagination: {} as Pagination,
      diploma: {} as Diploma,
      diplomaUser: {} ,
            };
  },
  actions: {
  
    setViewDiplomaDialog(value: boolean) {
      this.viewDiplomaDialog = value;
    },
    setVerifyDiplomaDialog(value: boolean) {
      this.verifyDiplomaDialog = value;
    },
    setAddDiplomaDialog(value: boolean,diplomaUser: Diploma) {
      this.diplomaUser = diplomaUser;
      this.addDiplomaDialog = value;
    },
    setViewDiplomaDialog(value: boolean, diploma: Diploma) {
      this.viewDiplomaDialog = value;
      this.diploma = diploma;
    },
  },
  getters: {
    getDiplomasColumns: (state) => state.diplomasColumns,
    getVerifyDiplomaDialog: (state) => state.verifyDiplomaDialog,
    getDiplomaUser: (state) => state.diplomaUser,
    getAddDiplomaDialog: (state) => state.addDiplomaDialog,
    getViewDiplomaDialog: (state) => state.viewDiplomaDialog,
    getDiploma: (state) => state.diploma,
  },
});
