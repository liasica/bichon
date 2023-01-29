export class ResponseError extends Error {
  public name: string = "ResponseError";
  public data;
  public code;
  constructor(message: string, code: number, data: any) {
    super(message);
    this.data = data;
    this.code = code;

    // Set the prototype explicitly.
    Object.setPrototypeOf(this, ResponseError.prototype);
  }
}
