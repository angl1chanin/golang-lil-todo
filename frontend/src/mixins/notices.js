import { createToast } from "mosha-vue-toastify";

export default {
    methods: {
        success(message) {
            createToast(message,{
                position: 'top-right',
                timeout: 2000,
                showIcon: true,
                hideProgressBar: true,
                type: 'success',
            })
        },
        danger(error) {
            createToast(`Error: ${error}`,{
                position: 'top-right',
                timeout: 2000,
                showIcon: true,
                hideProgressBar: true,
                type: 'danger',
            })
        },
        warning(message) {
            createToast(message,{
                position: 'top-right',
                timeout: 2000,
                showIcon: true,
                hideProgressBar: true,
                type: 'warning',
            })
        }
    }
}