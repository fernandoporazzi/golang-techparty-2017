import React, { Component } from 'react';
import './App.css';

class App extends Component {

  constructor() {
    super();

    this.state = {
      movies: []
    }
  }

  componentDidMount() {
    this.fetchData();
  }

  fetchData() {
    var self = this;

    fetch('http://localhost:8001/movie')
      .then((response) => {
        return response.json()
      })
      .then((data) => {
        self.setState({
          movies: data
        });
      })
      .catch((err) => {
        console.log(err);
      })
  }

  render() {
    const basePath = 'http://localhost:8001/uploads/' ;

    let m = this.state.movies.map((obj, index) => {
      return <span key={index}>{obj.name} <img src={basePath+obj.cover} /> </span>
    });

    return (
      <div className="App">
        {m}
      </div>
    );
  }
}

export default App;
