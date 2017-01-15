export class Article {
    constructor(
        public DateTime: string,
        public ID: string,
        public Title: string,
        public Author: string,
        public Content: string,
        public Published: boolean,
        public PictureFileName: string
    ) { }
}