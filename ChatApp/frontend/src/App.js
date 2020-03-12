// App.js
// Import our new component from it's relative path
import Header from './components/Header/Header';
import { Component } from 'react';
// ...
class App extends Component{
  render() {
    return (
      <div className="App">
        <Header />
        <button onClick={this.send}>Hit</button>
      </div>
    );
  }
}

export default App;
