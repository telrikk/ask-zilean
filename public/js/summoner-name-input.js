import React from 'react';
import $ from 'jquery';

export default class SummonerNameInput extends React.Component {
  constructor() {
    super();
    this.state = {
      name: '',
      buttonText: 'Search',
      errorText: '',
      isSearchComplete: false,
    };
    this.handleSubmit = this.handleSubmit.bind(this);
    this.handleTextChange = this.handleTextChange.bind(this);
  }
  handleTextChange(e) {
    this.setState({ name: e.target.value });
  }
  handleSubmit(e) {
    e.preventDefault();
    const name = this.state.name.trim();
    if (!name) {
      return;
    }
    this.setState({ errorText: '', buttonText: 'Loading...' });
    $.get({
      url: `/recentgames?name=${name}`,
      success: (data) => {
        this.props.onSuccess(data);
        this.setState({ isSearchComplete: true });
      },
      error: xhr => {
        if (xhr.status === 404) {
          this.setState({ errorText: `Could not find summoner with name ${name}` });
        } else {
          this.setState({ errorText: 'Internal server error. Please try again soon.' });
        }
        this.setState({ buttonText: 'Search' });
      },
    });
  }
  render() {
    let formGroupClass = 'form-group';
    if (this.state.errorText) {
      formGroupClass += ' has-error';
    }
    return (
      <form className={this.state.isSearchComplete ? 'hidden' : ''}>
        <div className={formGroupClass}>
          <div className="summoner-name-input input-group">
            <input
              type="text"
              placeholder="Summoner Name"
              className="form-control"
              onChange={this.handleTextChange}
            />
            <span className="input-group-btn">
              <button onClick={this.handleSubmit} type="submit" className="btn btn-primary">
                {this.state.buttonText}
              </button>
            </span>
          </div>
          <p className="help-block">
            {this.state.errorText}
          </p>
        </div>
      </form>
    );
  }
}

SummonerNameInput.propTypes = {
  onSuccess: React.PropTypes.func,
};
