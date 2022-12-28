import React, {Component} from "react";
import "./Orders.scss";
import OrderDetails from "../OrderDetails";

class Orders extends Component {
    render() {
        const orders = this.props.orders.map(order =>
            <OrderDetails key={order.id} order={order} />
        );

        return (
            <div className="Orders">
                <h2>Order # | Created At | Phone Number</h2>
                {orders}
            </div>
        )
    }
}

export default Orders;