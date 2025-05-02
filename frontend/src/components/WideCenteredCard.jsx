const WideCenteredCard = ({ children }) => {
  return (
    <div className="m-3 d-flex justify-content-center">
      <div
        className="card col-12 col-md-12 col-lg-12 col-xl-12"
        style={{ maxWidth: "2400px", borderRadius: "1rem" }}
      >
        <div className="card-body p-5">
          {children}
        </div>
      </div>
    </div>
  )
}
export default WideCenteredCard