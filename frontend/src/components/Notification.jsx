  const Notification = ({ message }) => {
    if (!message) {
      return null
    }

    return (
      <div className="row justify-content-center">
        <div className="col-12 col-md-8 col-lg-6 col-xl-5">
          <div className="card" style={{ borderRadius: "1rem" }}>
            <div className="card-body p-5">
              <h2 className="text-center">
                <p style={{ color: "green" }}>{message}</p>
              </h2>
            </div>
          </div>
        </div>
      </div>
  )
  }