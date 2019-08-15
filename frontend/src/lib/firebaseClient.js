export default class {
  setUser(user) {
    this.user = user;
  }

  async getToken() {
    if (!this.user) {
      return undefined;
    }
    return this.user.getIdToken();
  }
}
