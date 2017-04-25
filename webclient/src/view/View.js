import React, { Component } from 'react';
import Card from '../card/Card';


class View extends Component {
  render() {
    let movies = this.props.movies.map((obj, index) => {
      return <Card movie={obj} key={index} />
    });

    return (
      <section className="view">
        {movies}
      </section>
    );
  }
}

export default View;
