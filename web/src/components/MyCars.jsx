// create a new component called MyCars that will display a list of cars fetched from the api endpoint /api/mycars

import React from 'react';
import { Link } from 'react-router-dom';

function MyCars() {

    const [isLoading, setIsLoading] = React.useState(true);
    const [myCars, setMyCars] = React.useState([]);
    const [error, setError] = React.useState(null);
    
    React.useEffect(() => {
        fetch('/api/mycars')
        .then((res) => res.json())
        .then((data) => {
            setMyCars(data);
        })
        .catch((err) => {
            setError(err);
        })
        .finally(() => {
            setIsLoading(false);
        });
    });
    
    if (isLoading) return <p>Loading...</p>;
    if (error) return <p>Error: {error.message}</p>;
    
    return (
        <div>
        <h2>My Cars</h2>
        <ul>
            {myCars.map((car) => (
            <li key={car.id}>
                <Link to={`/car/${car.id}`}>{car.make} {car.model}</Link>
            </li>
            ))}
        </ul>
        </div>
    );
}

    export default MyCars;