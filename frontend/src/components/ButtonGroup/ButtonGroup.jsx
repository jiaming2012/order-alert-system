import React, {Component} from "react";
import "./ButtonGroup.scss";
import {updateOrderStatus} from "../../api/client";

class ButtonGroup extends Component {
    constructor(props) {
        super(props);

        let lifecycle;
        if (props.orderState === 'awaiting_pickup') {
            lifecycle = 'ALERT_SENT';
        } else {
            lifecycle = 'NOT_CLICKED';
        }

        this.state = {
            lifecycle: lifecycle
        }
    }

    deleteButtonClicked = () => {
        this.setState({ lifecycle: 'DELETE_BUTTON_CLICKED' });
    }

    confirmButtonClicked = () => {
        this.setState({ lifecycle: 'CONFIRM_BUTTON_CLICKED' });
    }

    submitButtonClicked = (msgType) => {
        switch (msgType) {
            case 'done': {
                updateOrderStatus(`${this.props.orderId}`, 'awaiting_pickup')
                    .then(res => {
                        this.setState({ lifecycle: 'ALERT_SENT' });
                    })
                    .catch(err => {
                        console.error(err);
                    })

                break;
            }

            case 'delete': {
                updateOrderStatus(`${this.props.orderId}`, 'deleted')
                    .catch(err => {
                        console.error(err);
                    })
                break;
            }

            case 'picked_up': {
                updateOrderStatus(`${this.props.orderId}`, 'picked_up')
                    .catch(err => {
                        console.error(err);
                    })
                break;
            }

            default: {
                console.error(`unknown msgType ${msgType} for submitButtonClicked`);
            }
        }
    }

    cancelButtonClicked = (state) => {
        switch (state) {
            case 'ALERT_SENT': {
                this.setState({ lifecycle: 'DELETE_BUTTON_CLICKED' });
                break;
            }

            case 'CONFIRM_BUTTON_CLICKED': {
                this.setState({ lifecycle: 'NOT_CLICKED' });
                break;
            }

            case 'delete': {
                updateOrderStatus(`${this.props.orderId}`, 'deleted')
                    .catch(err => {
                        console.error(err);
                    })
                break;
            }

            default: {
                console.error(`unknown currentState ${state} for cancelButtonClicked`);
            }
        }
    }

    render() {
        return (
            (this.state.lifecycle === 'CONFIRM_BUTTON_CLICKED') ? (
                <div className="ButtonGroup">
                    <button className="RedButton" onClick={this.cancelButtonClicked.bind(null, 'CONFIRM_BUTTON_CLICKED')}>
                        Cancel
                    </button>
                    <button className="GreenButton" onClick={this.submitButtonClicked.bind(null, 'done')}>
                        Notify
                    </button>
                </div>
            ) : ((this.state.lifecycle === 'DELETE_BUTTON_CLICKED') ? (
                    <div className="ButtonGroup">
                        <button className="RedButton" onClick={this.cancelButtonClicked.bind(null, 'CONFIRM_BUTTON_CLICKED')}>
                            Cancel
                        </button>
                        <button className="YellowButton" onClick={this.submitButtonClicked.bind(null, 'delete')}>
                            Delete
                        </button>
                    </div>
                ) : ((this.state.lifecycle === 'ALERT_SENT') ? (
                    <div className="ButtonGroup">
                        {/*<button className="YellowButton" onClick={this.cancelButtonClicked.bind(null, 'delete')}>*/}
                        {/*    Delete*/}
                        {/*</button>*/}
                        <button className="OrangeButton" onClick={this.submitButtonClicked.bind(null, 'picked_up')}>
                            Picked{'\u00A0'}Up
                        </button>
                    </div>
                ) : (
                <div className="ButtonGroupImage">
                    <button>
                        <img id="check-btn" src="/check.png" alt="confirm button" onClick={this.confirmButtonClicked}/>
                    </button>
                    <button>
                        <img id="delete-btn" src="/delete-button.png" alt="delete button" onClick={this.deleteButtonClicked}/>
                    </button>
                </div>
                ))
            )
        )
    }
}

export default ButtonGroup;