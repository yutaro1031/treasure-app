import { h, Component } from "preact";
import client from "../client";

class CreateForm extends Component {
  constructor() {
    super();
    this.state = {
      title: "",
      body: "",
      result: ""
    };

    this.onTitleChange = this.onTitleChange.bind(this);
    this.onBodyChange = this.onBodyChange.bind(this);
    this.onSubmit = this.onSubmit.bind(this);
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
      .postArticle(this.state.title, this.state.body)
      .then(v => this.setState({ result: `articleId: ${v.id}` }));
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
          />
          <br />
          <label htmlFor="bodyForm">body</label>
          <textarea
            id="bodyForm"
            rows={5}
            value={state.body}
            onChange={this.onBodyChange}
          />
          <br />
          <button onClick={this.onSubmit}>Submit</button>
        </form>
        <p>{String(state.result)}</p>
      </div>
    );
  }
}

export default CreateForm;
