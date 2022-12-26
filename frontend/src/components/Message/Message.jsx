import React, {Component} from "react";
import "./Message.scss";

class Message extends Component {
    constructor(props) {
        super(props);
        let temp = JSON.parse(this.props.message);
        console.log("temp: ", temp);
        this.state = {
            message: temp
        };
    }

    deleteButtonClicked() {
        console.log('delete button clicked!!')
    }

    confirmButtonClicked() {
        console.log('confirm button clicked!!');
    }

    padText(str) {
        const strLen = str.length;
        if (strLen > 10) {
            return str;
        }

        let pad = "";
        for (let i=0; i < 10-strLen; i++) {
            pad += "\u00A0";
        }

        return `${str}${pad}`
    }

    render() {
        return (
            <div className="Message">
                <div id="container">
                    <div className="containerText">
                        {this.padText(this.state.message.body)}
                    </div>
                    <div className="containerText">
                        856-875-7743
                    </div>
                    <div id="buttonContainer">
                        <button>
                            <img src="/check.png" alt="confirm button" onClick={this.confirmButtonClicked}/>
                        </button>
                        <button>
                            <img src="/delete-button.png" alt="delete button" onClick={this.deleteButtonClicked}/>
                        </button>
                    </div>
                </div>
            </div>
        )
    }
}

export default Message;