import { useEffect, useState } from "react";
import PropTypes from "prop-types";

function Car({ id }) {
  const [car, setCar] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  useEffect(() => {
    fetch(`/api/carWithServiceRecord/${id}`)
      .then((res) => res.json())
      .then((data) => {
        setCar(data);
      })
      .catch((err) => {
        setError("There was a problem fetching the car");
        console.error(err);
      })
      .finally(() => {
        setLoading(false);
      });
  }, [id]);

  if (loading) return <p>Loading...</p>;

  if (error) return <p>Error: {error}</p>;

  return (
    <div>
      <h1>
        {car.make} {car.model}
      </h1>
    </div>
  );
}
Car.propTypes = {
  id: PropTypes.string.isRequired,
};

export default Car;
