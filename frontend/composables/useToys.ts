export const useToys = () => {
    const toys = ref([
        {
            id: 1,
            title: "LEGO City Police Station",
            description: "Полный набор со всеми минифигурками, включая полицейскую машину и вертолет",
            age: "6-12y",
            location: "Москва",
            condition: "Отличное",
            image: "toy1.jpg"
        },
        {
            id: 2,
            title: "Barbie Dreamhouse",
            description: "Трехэтажный дом с мебелью и аксессуарами",
            age: "3-10y",
            location: "Санкт-Петербург",
            condition: "Хорошее",
            image: "toy2.jpg"
        },
        {
            id: 3,
            title: "Hot Wheels Track Set",
            description: "Огромной набор трасс с петлями и трамплинами, включает 5 машинок",
            age: "5-12y",
            location: "Казань",
            condition: "Новое",
            image: "toy3.jpg"
        },
        {
            id: 4,
            title: "Wooden Train Set",
            description: "Классическая деревянная железная дорога с мостами и станциями",
            age: "2-8y",
            location: "Новосибирск",
            condition: "Хорошее",
            image: "toy4.jpg"
        },
        {
            id: 5,
            title: "Play-Doh Kitchen Set",
            description: "Полный кухонный набор с формочками и инструментами",
            age: "3-8y",
            location: "Екатеринбург",
            condition: "Отличное",
            image: "toy5.jpg"
        },
        {
            id: 6,
            title: "Marvel Action Figures",
            description: "Коллекция из 6 фигурок Мстителей",
            age: "4-12y",
            location: "Москва",
            condition: "Хорошее",
            image: "toy6.jpg"
        },
        {
            id: 7,
            title: "Baby Activity Gym",
            description: "Красочный игровой коврик с подвесными игрушками",
            age: "0-12m",
            location: "Санкт-Петербург",
            condition: "Отличное",
            image: "toy7.jpg"
        },
        {
            id: 8,
            title: "Nintendo Switch",
            description: "Консоль с двумя контроллерами и 3 играми",
            age: "6+",
            location: "Казань",
            condition: "Хорошее",
            image: "toy8.jpg"
        },
        {
            id: 9,
            title: "Science Kit",
            description: "Образовательный набор для экспериментов с микроскопом",
            age: "8-14y",
            location: "Новосибирск",
            condition: "Новое",
            image: "toy9.jpg"
        },
        {
            id: 10,
            title: "Dollhouse Furniture",
            description: "Полный набор миниатюрной мебели",
            age: "4-10y",
            location: "Екатеринбург",
            condition: "Отличное",
            image: "toy10.jpg"
        },
        {
            id: 11,
            title: "Remote Control Car",
            description: "Внедорожник 4x4 на радиоуправлении с пультом",
            age: "8-14y",
            location: "Москва",
            condition: "Хорошее",
            image: "toy11.jpg"
        },
        {
            id: 12,
            title: "Art Supply Set",
            description: "Полный художественный набор с мольбертом и принадлежностями",
            age: "5+",
            location: "Санкт-Петербург",
            condition: "Новое",
            image: "toy12.jpg"
        },
        {
            id: 13,
            title: "Board Game Collection",
            description: "Набор из 5 популярных настольных игр",
            age: "6+",
            location: "Казань",
            condition: "Отличное",
            image: "toy13.jpg"
        },
        {
            id: 14,
            title: "Musical Keyboard",
            description: "Электронное пианино с обучающим режимом",
            age: "5-12y",
            location: "Новосибирск",
            condition: "Хорошее",
            image: "toy14.jpg"
        },
        {
            id: 15,
            title: "Building Blocks Set",
            description: "Более 200 деталей с контейнером для хранения",
            age: "3-8y",
            location: "Екатеринбург",
            condition: "Отличное",
            image: "toy15.jpg"
        }
    ])

    return {
        toys
    }
}