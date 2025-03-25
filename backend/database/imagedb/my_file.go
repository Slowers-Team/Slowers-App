package cloudfor

// Import Cloudinary and other necessary libraries
//===================
import (
	"context"
	"fmt"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api"
	"github.com/cloudinary/cloudinary-go/v2/api/admin"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func credentials() (*cloudinary.Cloudinary, context.Context) {
	// Add your Cloudinary credentials, set configuration parameter
	// Secure=true to return "https" URLs, and create a context
	//===================
	cld, _ := cloudinary.New()
	cld.Config.URL.Secure = true
	ctx := context.Background()
	return cld, ctx
}

func uploadImage(cld *cloudinary.Cloudinary, ctx context.Context) {

	// Upload the image.
	// Set the asset's public ID and allow overwriting the asset with new versions
	resp, err := cld.Upload.Upload(ctx, "https://cloudinary-devs.github.io/cld-docs-assets/assets/images/butterfly.jpeg", uploader.UploadParams{
		PublicID:       "quickstart_butterfly",
		UniqueFilename: api.Bool(false),
		Overwrite:      api.Bool(true)})
	if err != nil {
		fmt.Println("error")
	}

	// Log the delivery URL
	fmt.Println("****2. Upload an image****\nDelivery URL:", resp.SecureURL)
}

func getAssetInfo(cld *cloudinary.Cloudinary, ctx context.Context) {
	// Get and use details of the image
	// ==============================
	resp, err := cld.Admin.Asset(ctx, admin.AssetParams{PublicID: "quickstart_butterfly"})
	if err != nil {
		fmt.Println("error")
	}
	fmt.Println("****3. Get and use details of the image****\nDetailed response:\n", resp)
	// Assign tags to the uploaded image based on its width. Save the response to the update in the variable 'update_resp'.
	if resp.Width > 900 {
		update_resp, err := cld.Admin.UpdateAsset(ctx, admin.UpdateAssetParams{
			PublicID: "quickstart_butterfly",
			Tags:     []string{"large"}})
		if err != nil {
			fmt.Println("error")
		} else {
			// Log the new tag to the console.
			fmt.Println("New tag: ", update_resp.Tags)
		}
	} else {
		update_resp, err := cld.Admin.UpdateAsset(ctx, admin.UpdateAssetParams{
			PublicID: "quickstart_butterfly",
			Tags:     []string{"small"}})
		if err != nil {
			fmt.Println("error")
		} else {
			// Log the new tag to the console.
			fmt.Println("New tag: ", update_resp.Tags)
		}
	}

}

func transformImage(cld *cloudinary.Cloudinary, ctx context.Context) {
	// Instantiate an object for the asset with public ID "my_image"
	qs_img, err := cld.Image("quickstart_butterfly")
	if err != nil {
		fmt.Println("error")
	}

	// Add the transformation
	qs_img.Transformation = "r_max/e_sepia"

	// Generate and log the delivery URL
	new_url, err := qs_img.String()
	if err != nil {
		fmt.Println("error")
	} else {
		print("****4. Transform the image****\nTransfrmation URL: ", new_url)
	}
}

func main() {
	cld, ctx := credentials()
	uploadImage(cld, ctx)
	getAssetInfo(cld, ctx)
}
