import './App.css';
import { connect } from "./api";
import Header from "./components/Header/Header";
import {Component} from "react";
import Orders from "./components/Orders";

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
        orders: []
    }
  }

  componentDidMount() {
      connect((orders) => {
          this.setState(prevState => ({
              orders
          }))
      })
  }

  render() {
    return (
        <div className="App">
          <Header />
            <Orders orders={this.state.orders} />
        </div>
    );
  }
}

export default App;
