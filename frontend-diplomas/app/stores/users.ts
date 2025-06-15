import { defineStore } from "pinia";


const usersColumns = [


  {
    name: "firstName",
    required: true,
    label: "First Name",
    align: "left",
    sortable: true,
  },    
  { name: "lastName", label: "Last Name", field: "lastName", align: "center", sortable: true },
  {
    name: "email",
    label: "Email",
    field: "email",
    align: "center",
    sortable: true,
  },
  {
    name: "phone",
    label: "Phone",
    field: "phone",
    align: "center",
    sortable: true,
  },
  {
    name: "university",
    label: "University",
    field: "university",
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



export const useUsersStore = defineStore("usersStore", {
  state: () => {
    return {
      perPage: 20,
      users: [] as Array<any>,
      usersColumns: usersColumns,
      addUserDialog:false,
      viewUserDialog:false,
      pagination: {} as Pagination,
      user: {} as User,
      addUserDiplomaDialog:false,
      };
  },
  actions: {
    async setUsers(users: Array<any > ) {  
      this.users = users; 
    },

    setAddUserDialog(value: boolean) {
      this.addUserDialog = value;
    },
    setViewUserDialog(value: boolean, user: User) {
      this.viewUserDialog = value;
      this.user = user;   
    },

    setAddUserDiplomaDialog(value: boolean, user: User) {
      this.addUserDiplomaDialog = value;
      this.user = user;   
    },
  },
  getters: {
    getUsers: (state) => state.users,
    getUsersColumns: (state) => state.usersColumns,
    getAddUserDialog: (state) => state.addUserDialog,
    getViewUserDialog: (state) => state.viewUserDialog,
    getUser: (state) => state.user,
    getAddUserDiplomaDialog: (state) => state.addUserDiplomaDialog,
  },
});
