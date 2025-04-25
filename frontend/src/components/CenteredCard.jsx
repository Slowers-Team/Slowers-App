const CenteredCard = ({ children }) => {
  return (
    <div className="m-3 d-flex justify-content-center">
      <div
        className="card col-12 col-md-10 col-lg-10 col-xl-8"
        style={{ maxWidth: "1400px", borderRadius: "1rem" }}
      >
        <div className="card-body p-5">
          {children}
        </div>
      </div>
    </div>
  )
}
export default CenteredCard