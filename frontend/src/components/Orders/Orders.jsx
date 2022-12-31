import React, {Component} from "react";
import "./Orders.scss";
import OrderDetails from "../OrderDetails";

class Orders extends Component {
    render() {
        const orders = this.props.orders.length > 0 ? (
            this.props.orders.map(order =>
                <OrderDetails key={order.ID} order={order} />
            )
        ) : null;

        console.log('props: ', this.props)

        return (
            <div className="Orders">
                <table>
                    <thead>
                        <tr>
                            <th>Order #</th>
                            <th>Created At</th>
                            <th>Phone Number</th>
                            <th></th>
                        </tr>
                    </thead>
                    <tbody>
                        {orders}
                    </tbody>
                </table>
            </div>
        )
    }
}

export default Orders;