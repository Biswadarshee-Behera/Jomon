import axios from "axios";
export const applicationDetail = {
  state: {
    application_id: "",
    created_at: "",
    applicant: {
      trap_id: "",
      is_admin: false
    },
    current_detail: {
      update_user: {
        trap_id: "",
        is_admin: ""
      },
      type: "",
      title: "",
      remarks: "",
      ammount: 0,
      paid_at: "",
      updated_at: "",
      repaid_to_id: []
    },
    current_state: "",
    images: [""],
    comments: [
      {
        comment_id: 0,
        user: {
          trap_id: "",
          is_admin: false
        },
        comment: "",
        created_at: "",
        updated_at: ""
      }
    ],
    state_logs: [
      {
        update_user: {
          trap_id: "",
          is_admin: false
        },
        to_state: "",
        reason: "",
        created_at: ""
      }
    ],
    application_detail_logs: [
      {
        update_user: {
          trap_id: "",
          is_admin: false
        },
        type: "",
        title: "",
        remarks: "",
        ammount: 0,
        paid_at: "",
        updated_at: ""
      }
    ],
    repayment_logs: [
      {
        repaid_by_user: {
          trap_id: "",
          is_admin: false
        },
        repaid_to_user: {
          trap_id: "",
          is_admin: false
        },
        repaid_at: ""
      }
    ]
  },
  getters: {
    logs: state => {
      let logs = [];
      state.comments.forEach(log => {
        logs.push({ log_type: "comment", content: log });
      });
      state.state_logs.forEach(log => {
        logs.push({ log_type: "state", content: log });
      });
      state.application_detail_logs.forEach(log => {
        logs.push({ log_type: "application", content: log });
      });
      state.repayment_logs.forEach(log => {
        logs.push({ log_type: "repayment", content: log });
      });
      //TODO:時系列順にlogを並び変える
      return logs;
    }
  },
  mutations: {
    setApplicationDetail(state, newState) {
      Object.keys(state).forEach(key => {
        state[key] = newState[key];
      });
    }
  },
  actions: {
    async getApplicationDetail({ commit }, applicationId) {
      try {
        const response = await axios.get("/api/applications/" + applicationId);
        commit("setApplicationDetail", response.data);
      } catch (err) {
        console.log(err);
      }
    }
  }
};
