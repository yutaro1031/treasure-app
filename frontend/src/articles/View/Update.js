import { h, Component } from "preact";
import client from "../client";

class CreateForm extends Component {
  constructor({ articleId }) {
    super();
    this.state = {
      title: "",
      body: "",
      result: "",
      initialized: false
    };

    this.onTitleChange = this.onTitleChange.bind(this);
    this.onBodyChange = this.onBodyChange.bind(this);
    this.onSubmit = this.onSubmit.bind(this);
  }

  componentDidMount() {
    client.fetchArticle(this.props.articleId).then(v => {
      this.setState({ title: v.title, body: v.body, initialized: true });
    });
  }

  onTitleChange(e) {
    this.setState({ title: e.target.value });
  }

  onBodyChange(e) {
    this.setState({ body: e.target.value });
  }

  onSubmit(e) {
    e.preventDefault();
    client
      .updateArticle(this.props.articleId, this.state.title, this.state.body)
      .then(v => this.setState({ result: `更新が完了しました` }));
  }

  render(_, state) {
    return (
      <div>
        <form>
          <label htmlFor="titleForm">Title</label>
          <input
            id="titleForm"
            type="text"
            value={state.title}
            onChange={this.onTitleChange}
            disabled={!state.initialized}
          />
          <br />
          <label htmlFor="bodyForm">body</label>
          <textarea
            id="bodyForm"
            rows={5}
            value={state.body}
            onChange={this.onBodyChange}
            disabled={!state.initialized}
          />
          <br />
          <button onClick={this.onSubmit} disabled={!state.initialized}>
            Submit
          </button>
        </form>
        <p>{String(state.result)}</p>
      </div>
    );
  }
}

export default CreateForm;
