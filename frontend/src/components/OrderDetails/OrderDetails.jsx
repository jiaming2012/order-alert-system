import React, {Component} from "react";
import "./OrderDetails.scss";
import ButtonGroup from "../ButtonGroup/ButtonGroup";
import moment from "moment";

class OrderDetails extends Component {
    orderId;
    createdAt;
    phoneNumber;
    status;

    constructor(props) {
        super(props);
    }

    render() {
        let createdAt = moment(this.props.order.createdAt).format("h:mm a");

        return (
            <div className="Message">
                <div id="container">
                    <div className="containerText">
                        {this.props.order.orderId}
                    </div>
                    <div className="containerText">
                        {createdAt}
                    </div>
                    <div className="containerText">
                        {this.props.order.phoneNumber}
                    </div>
                    <ButtonGroup orderState={this.props.order.status} />
                </div>
            </div>
        )
    }
}

export default OrderDetails;