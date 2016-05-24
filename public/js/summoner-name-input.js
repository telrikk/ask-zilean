import React from 'react';
import $ from 'jquery';

export default class SummonerNameInput extends React.Component {
  constructor() {
    super();
    this.state = { name: '' };
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
    $.get({
      url: `/recentgames?name=${name}`,
      success: this.props.onSuccess,
      error(xhr) {
        if (xhr.status === 404) {
          $('.form-group').addClass('has-error');
          $('.help-block').html(`Could not find summoner with name ${name}`);
        } else {
          $('.form-group').addClass('has-error');
          $('.help-block').html('Internal server error. Please try again soon.');
        }
      },
    });
  }
  render() {
    return (
      <form>
        <div className="form-group">
          <div className="summoner-name-input input-group">
            <input
              type="text"
              placeholder="Summoner Name"
              className="form-control"
              onChange={this.handleTextChange}
            />
            <span className="input-group-btn">
              <button onClick={this.handleSubmit} type="submit" className="btn btn-primary">
                Search
              </button>
            </span>
          </div>
          <p className="help-block" />
        </div>
      </form>
    );
  }
}

SummonerNameInput.propTypes = {
  onSuccess: React.PropTypes.func,
};
