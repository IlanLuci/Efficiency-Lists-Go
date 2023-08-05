class Storage
{
    // Storage should only be initialized once for the program.
    // more than one instance of this class is not neccesary
    // and may create problems

    data = {};
    lists = [
        'general',
        'today',
        'upcoming',
    ];

    constructor(saved)
    {
        const id = localStorage.getItem('id');

        if (saved && Object.keys(saved).length > 0)
        {
            // list exists in storage from previous use of the website
            // load this list as our basis
            this.data = saved;

            for (let topic in this.data.topics)
            {
                this.lists.push(topic)
            }
        }
        else
        {
            // add basic lists to the storage
            console.log('initializing lists');

            for (let list of this.lists)
            {
                this.data[list] = [];
            }

            this.data.topics = {};

            // by default set general as selected list
            this.data.selected = 'general';
        }

        window.addEventListener('beforeunload', async () => 
        {
            // save updated list back to storage
            console.log('saving lists to storage...');

            await fetch(`/api/v1/data?id=${id}`, { method: 'post', body: JSON.stringify(this.data), headers: { 'Content-Type': 'application/json' } });
        });
    }

    addItem(list, name, date, description)
    {
        if (storage.data.hasOwnProperty(list))
        {
            // this is a basic built in list, add to it
            storage.data[list].push({"name": name, "checked": false, "date": date, "description": description || ''});
        }
        else
        {
            // this is a user added list
            storage.data.topics[list].push({"name": name, "checked": false, "date": date, "description": description || ''});
        }
    }
}