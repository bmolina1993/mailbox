export const useFetch = async () => {
  const response = await fetch("http://localhost:8080/");
  const auxData = await response.json();

  //formateo a estructura deseada
  const data = auxData.map((item) => {
    return {
      index: item._index,
      body: item._source.Body,
      date: item._source.Date,
      folder: item._source.Folder,
      from: item._source.From,
      subject: item._source.Subject,
      to: item._source.To,
    };
  });
  return data;
};

//obtiene usuarios random par usar imagen perfil
export const fetchRandomUser = async () => {
  const response = await fetch("https://randomuser.me/api/?results=150", {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
  });

  const { results } = await response.json();

  const data = results.map((item) => {
    return {
      picture: item.picture.thumbnail,
      name: null,
    };
  });

  return data;
};

// format ["Wed, 21 Nov 2001 05:13:17 -0800 (PST)"]
// to custom data defined
export const getDate = (auxDate, type) => {
  const date = new Date(auxDate);

  switch (type) {
    case "mmm":
      return new Intl.DateTimeFormat("es-ar", {
        month: "short",
      }).format(new Date(date));

    case "dd":
      return new Date(date).getDay();

    case "long":
      return new Date(date).toLocaleString("es-CO");

    default:
      return new Date(date);
  }
};

export const getFirstLetter = (word = "") => {
  return word.slice(0, 1).toLocaleUpperCase();
};

export const proxyToObject = (data) => JSON.parse(JSON.stringify(data));
