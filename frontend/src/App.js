import { h, Component } from "preact";
import firebase from "./firebase";
import { getPrivateMessage, getPublicMessage } from "./api";
import ArticleCreateView from "./articles/View/Create";
import ArticleShowView from "./articles/View/Show";
import ArticleUpdateView from "./articles/View/Update";

import ArticleClient from "./articles/client";

class App extends Component {
  constructor() {
    super();
    this.state.user = null;
    this.state.message = "";
    this.state.errorMessage = "";
  }

  initializeClients(user) {
    [ArticleClient].map(v => v.setUser(user));
  }

  componentDidMount() {
    firebase.auth().onAuthStateChanged(user => {
      if (user) {
        this.setState({ user });
        this.initializeClients(user);
      } else {
        this.initializeClients(null);
        this.setState({
          user: null
        });
      }
    });
  }

  getPrivateMessage() {
    this.state.user
      .getIdToken()
      .then(token => {
        return getPrivateMessage(token);
      })
      .then(resp => {
        this.setState({
          message: resp.message
        });
      })
      .catch(error => {
        this.setState({
          errorMessage: error.toString()
        });
      });
  }

  render(props, state) {
    if (state.user === null) {
      return <button onClick={firebase.login}>Please login</button>;
    }
    return (
      <div>
        <div>{state.message}</div>
        <p style="color:red;">{state.errorMessage}</p>
        <button onClick={this.getPrivateMessage.bind(this)}>
          Get Private Message
        </button>
        <button onClick={firebase.logout}>Logout</button>

        <ArticleCreateView />
        <ArticleShowView articleId={1} />
        <ArticleUpdateView articleId={3} />

        <button
          onClick={_ => {
            ArticleClient.deleteArticle(2).then(v => console.log(v));
          }}
        >
          消す
        </button>
      </div>
    );
  }
}

export default App;
