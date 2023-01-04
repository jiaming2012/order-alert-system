import axios from "axios";

const backendURL = process.env['REACT_APP_BACKEND_URL'];
const basicAuthUser = process.env['REACT_APP_BASIC_AUTH_USER'];
const basicAuthPass = process.env['REACT_APP_BASIC_AUTH_PASS'];

/**
 *
 * @param id - the backend id of the order
 * @param status - the new status that the order will be updated to
 */
function updateOrderStatus(id, status) {
    return axios.post(`${backendURL}/admin/order`, {
        id,
        status
    }, { auth: {
            username: basicAuthUser,
            password: basicAuthPass
        }})
}

export { updateOrderStatus };
