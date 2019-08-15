import { h, Component } from "preact";
import client from "../client";

export default class extends Component {
  constructor({ articleId }) {
    super();
    this.state = {
      article: undefined
    };
  }

  componentDidMount() {
    client
      .fetchArticle(this.props.articleId)
      .then(v => this.setState({ article: v }));
  }

  render(_, state) {
    if (state.article === undefined) {
      return <div>読込中</div>;
    }

    return (
      <ArticleRender title={state.article.title} body={state.article.body} />
    );
  }
}

const ArticleRender = ({ title, body }) => (
  <div>
    <p>{title}</p>
    <hr />
    <div style={{ border: "solid 1px black", padding: "8px" }}>{body}</div>
  </div>
);
