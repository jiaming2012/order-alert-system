import React, {Component} from "react";
import "./Message.scss";
import ButtonGroup from "../ButtonGroup/ButtonGroup";

class Message extends Component {
    constructor(props) {
        super(props);
        let temp = JSON.parse(this.props.message);
        console.log("temp: ", temp);
        this.state = {
            message: temp
        };
    }

    render() {
        return (
            <div className="Message">
                <div id="container">
                    <div className="containerText">
                        {this.state.message.body}
                    </div>
                    <div className="containerText">
                        856-875-7743
                    </div>
                    <ButtonGroup />
                </div>
            </div>
        )
    }
}

export default Message;