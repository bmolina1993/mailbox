export const useFetch = async () => {
  const response = await fetch("http://localhost:4080/es/_search", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Authorization: "Basic " + btoa("admin:Complexpass#123"),
    },
    body: JSON.stringify({
      search_type: "matchall",
      from: 0,
      size: 517424,
    }),
  });

  //se deja alias [data] a campo [hits]
  const {
    hits: { hits: data },
  } = await response.json();

  /*
  //[arora-h] -> usuario con sub-carpetas
  const dataV2 = data.filter(
    (item) => (item._index == "arora-h") | (item._index == "allen-p")
  );
  //console.log("dataV2:", dataV2);

  // constante con data final de estructura deseada
  // filtra por carpeta [inbox] que contiene sub-carpetas
  const dataV3 = dataV2
    .map((item) => {
      return {
        index: item._index,
        body: item._source.Body,
        date: item._source.Date,
        folder: item._source.Folder,
        from: item._source.From,
        subject: item._source.Subject,
        to: item._source.To,
      };
    })
    .filter((item) =>
      ["deleted_items", "inbox"].some((itemCond) =>
        item.folder.includes(itemCond)
      )
    );

  return dataV3;
  */

  //[FINAL]
  const dataReturn = data.map((item) => {
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
  return dataReturn;
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
  return word.substr(0, 1).toLocaleUpperCase();
};

export const proxyToObject = (data) => JSON.parse(JSON.stringify(data));

/*

  //obtiene lista de usuarios
  //const auxUserFolder = dataV3.map((item) => item.index);
  const auxUserFolder = dataV3.map((item) => {
    return {
      index: item.index,
      folder: item.folder,
    };
  });
  console.log("auxUserFolder:", auxUserFolder);

  const auxUser = dataV3.map((item) => {
    return {
      index: item.index,
    };
  });
  console.log("auxUser:", auxUser);

  //quita duplicados
  //const userFolder = [...new Set(auxUserFolder)];
  //console.log("userFolder:", userFolder);

  let user = auxUser.reduce((prevValue, currentValue) => {
    if (!prevValue.some((obj) => obj.index === currentValue.index)) {
      prevValue.push(currentValue);
    }
    return prevValue;
  }, []);
  console.log("user:", user);

  let userFolder = auxUserFolder.reduce((prevValue, currentValue) => {
    if (
      !prevValue.some(
        (obj) =>
          obj.index === currentValue.index && obj.folder === currentValue.folder
      )
    ) {
      prevValue.push(currentValue);
    }
    return prevValue;
  }, []);
  console.log("userFolder:", userFolder);

  //separa en grupo de usuario
  const arrPerUser = [];

  for (let idx in userFolder) {
    //obtiene todo elem. de dato por indice
    const elem = userFolder[idx];

    // if (arrPerUser.index == elem.index) {
    //   arrPerUser.push(elem.folder);
    // }

    user.map((item, idx) => {
      if (item.index == elem.index) {
        //arrPerUser[idx] = [];
        arrPerUser[idx].push(elem);
      }
    });
  }

  console.log("arrPerUser:", arrPerUser);

  //objetivo a lograr
  const objToCompare = {
    index: "arora-h",
    folder: "inbox",
    items: [],
    children: [
      {
        folder: "inbox/saved_mail",
        items: [],
        children: [],
      },
      {
        folder: "inbox/hist_vols",
        items: [],
        children: [
          {
            folder: "inbox/hist_vols/otro",
            items: [],
            children: [],
          },
        ],
      },
    ],
  };

  //agrega elems.hijos a cada sub-carpeta de json anterior
  for (let key in dataV3) {
    //obtiene todo el indice de dato
    const elem = dataV3[key];
    //si carpeta raiz concide con array de datos, ingresa los que sean iguales [1/6]
    if (objToCompare.folder == elem.folder) {
      objToCompare.items.push(elem);
    }

    //a la sub-carpeta del root, agrega los elementos [2/6]
    objToCompare.children.map((item, idx) => {
      if (item.folder == elem.folder) {
        objToCompare.children[idx].items.push(elem);
      }
    });

    //console.log("elem:", elem);
  }
  console.log("objToCompare:", objToCompare);
  */
