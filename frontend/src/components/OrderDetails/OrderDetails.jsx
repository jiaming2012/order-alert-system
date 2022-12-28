import React, {Component} from "react";
import "./Message.scss";
import ButtonGroup from "../ButtonGroup/ButtonGroup";

class Message extends Component {
    orderId;
    phoneNumber;
    status;

    constructor(props) {
        super(props);
    }

    render() {
        return (
            <div className="Message">
                <div id="container">
                    <div className="containerText">
                        {this.props.order.orderId}
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

export default Message;