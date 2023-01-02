import React, {Component} from "react";
import "./OrderDetails.scss";
import ButtonGroup from "../ButtonGroup/ButtonGroup";
import moment from "moment";

class OrderDetails extends Component {
    orderNumber;
    createdAt;
    phoneNumber;
    status;

    render() {
        let createdAt = moment(this.props.order.createdAt).format("h:mm a");

        return (
            <tr>
                <td>{this.props.order.orderNumber}</td>
                <td>{createdAt}</td>
                <td>{this.props.order.phoneNumber}</td>
                <td><ButtonGroup orderId={this.props.order.ID} orderState={this.props.order.status} /></td>
            </tr>
        )
    }
}

export default OrderDetails;