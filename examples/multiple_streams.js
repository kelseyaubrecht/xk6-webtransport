import wt from "k6/x/webtransport";

export default function () {
  const url = "https://localhost:443/webtransport";
  let data = [0, 1, 2, 3, 4];

  wt.connect(url);

  let streamOneid = wt.openStream();
  wt.write(data);

  let streamTwoid = wt.openStream();
  wt.write(data);

  wt.setActiveStream(streamOneid);
  wt.write(data);

  wt.close();
}
