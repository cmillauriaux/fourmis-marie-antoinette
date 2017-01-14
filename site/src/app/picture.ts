export class Picture {
    constructor(
        public FileName: string,
        public Link: string,
        public DateTime: number,
        public Next: number,
        public CameraID: number
    ) { }
}