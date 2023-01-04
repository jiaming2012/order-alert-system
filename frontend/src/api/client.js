import axios from "axios";

/**
 *
 * @param id - the backend id of the order
 * @param status - the new status that the order will be updated to
 */
function updateOrderStatus(id, status) {
    // todo: add a server login method (maybe cookies? environment variables are injected into REACT apps at build time!)
    return axios.post('/admin/order', {
        id,
        status
    })
}

export { updateOrderStatus };
