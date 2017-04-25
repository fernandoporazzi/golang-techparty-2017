import React, { Component } from 'react';
import {cdnPath} from '../util/config';
import './card.css';

class Card extends Component {
  render() {
    let m = this.props.movie;

    let categories = m.categories.map((obj, index) => {
      return <span key={index}>{obj}</span>
    })

    return (
      <div className="card">
        <img src={cdnPath+m.cover} alt={m.name} />
        <h4>{m.name}</h4>
        <h5>{m.year}</h5>

        <div className="card-categories">
          {categories}
        </div>

        <p>{m.synopsis}</p>
      </div>
    );
  }
}

export default Card;
