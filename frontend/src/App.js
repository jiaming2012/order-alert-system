import './App.css';
import { connect, sendMsg } from "./api";
import Header from "./components/Header/Header";
import {Component} from "react";
import Orders from "./components/OrdersHistory";

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
        orders: []
    }
    connect();
  }

  componentDidMount() {
      connect((order) => {
          console.log("New Message");
          this.setState(prevState => ({
              orders: [...this.state.orders, order]
          }))
          console.log(this.state);
      })
  }

  send() {
    console.log("hello");
    sendMsg("hello");
  }

  render() {
    return (
        <div className="App">
          <Header />
            <Orders orders={this.state.orders} />
            <button onClick={this.send}>Hit</button>
        </div>
    );
  }
}

export default App;
