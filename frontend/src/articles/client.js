import FirebaseClient from "../lib/firebaseClient";

const API_ENDPOINT = process.env.BACKEND_API_BASE;

class ArticleClient extends FirebaseClient {
  async postArticle(title, body) {
    const token = await this.getToken();
    return fetch(`${API_ENDPOINT}/articles`, {
      method: "POST",
      headers: {
        Authorization: `Bearer ${token}`
      },
      body: JSON.stringify({ title, body })
    }).then(v => v.json());
  }

  async fetchArticle(articleId) {
    return fetch(`${API_ENDPOINT}/articles/${articleId}`, {
      method: "GET"
    }).then(v => v.json());
  }

  async updateArticle(articleId, title, body) {
    const token = await this.getToken();
    return fetch(`${API_ENDPOINT}/articles/${articleId}`, {
      method: "PUT",
      headers: {
        Authorization: `Bearer ${token}`
      },
      body: JSON.stringify({ title, body })
    }).then(v => v.json());
  }

  async deleteArticle(articleId) {
    const token = await this.getToken();
    return fetch(`${API_ENDPOINT}/articles/${articleId}`, {
      method: "DELETE",
      headers: {
        Authorization: `Bearer ${token}`
      },
      body: JSON.stringify({})
    }).then(v => v.text());
  }
}

export default new ArticleClient();
