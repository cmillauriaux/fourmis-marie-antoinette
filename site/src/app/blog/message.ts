export class Message {
    constructor(
        public SignInURL: string,
        public SignOutURL: string,
        public NeedAdminAuthorization: boolean,
        public IsLogin: boolean,
        public IsAdmin: boolean,
        public Error: boolean,
        public ErrorMessage: string
    ) { }
}