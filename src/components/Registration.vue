<!-- Registration.vue -->
<template>
  <Vueform v-if="response.status == -1" v-bind="vueform" @success="handleResponse" @error="handleError" />
  <div v-else-if="response.status == 200" class="form-submitted">
    <h1>Successfully Registered!</h1>
    <p>Thank you for registering for HackNJIT 2025. We'll be in touch soon!</p>
    <RouterLink to="/" class="nav-link">
      <button class="vf-btn vf-btn-primary">Return Home</button>
    </RouterLink>
  </div>
  <div v-else class="form-submitted">
    <h1>Failed to register!</h1>
    <p>We could not register you for HackNJIT 2025 at this time. We may be having some technical difficulties. Check
      back in later.</p>
    <RouterLink to="/" class="nav-link">
      <button class="vf-btn vf-btn-primary">Return Home</button>
    </RouterLink>
  </div>
</template>


<script>
import { useVueform, Vueform } from '@vueform/vueform'
import schoolsJson from '../data/schools.json'
import countriesJson from '../data/countries.json'

export default {
  mixins: [Vueform],
  setup: useVueform,
  data: () => ({
    response: {
      status: -1,
    },
    vueform: {
      size: 'md',
      displayErrors: true,
      endpoint: '/api/register',
      action: '/api/register',
      method: 'POST',
      uploadFileOnChange: false, // Disable automatic file (PDF) uploads
      uploadTempFile: false,
      enctype: 'multipart/form-data', // For PDF (resume) uploads
      validateOn: 'step|change',
      steps: {
        page0: {
          label: 'Personal Information',
          elements: [
            'HackNJIT Registration',
            'divider',
            'container',
            'preferred_name_container',
            'email',
            'phone',
            'age',
            'country',
            'uni',
            'major',
            'otherMajor',
            'container2',
            'container3',
            'divider_1',
          ],
        },
        page1: {
          label: 'MLH Agreements',
          elements: [
            'h1',
            'divider_4',
            'mlh_checkbox_0',
            'mlh_checkbox_1',
            'mlh_checkbox_2',
            'divider_3',
          ],
        },
        page2: {
          elements: [
            'h1_2',
            'p_1',
            'minority',
            'gender',
            'race',
            'divider_6',
            'h1_1',
            'p',
            'divider_2',
            'resume',
            'linkedin',
            'divider_5',
          ],
          label: 'Extra Information',
        },
      },
      addClass: 'vf-registration',
      schema: {
        'HackNJIT Registration': {
          type: 'static',
          content: 'HackNJIT Registration',
          tag: 'h1',
        },
        divider: {
          type: 'static',
          tag: 'hr',
        },
        container: {
          type: 'group',
          schema: {
            first_name: {
              label: "First Name",
              type: 'text',
              columns: {
                sm: {
                  container: 12,
                  label: 12,
                  wrapper: 12,
                },
                lg: {
                  container: 6,
                  label: 12,
                  wrapper: 12,
                },
              },
              fieldName: 'First name',
              rules: [
                'required',
                'max:255',
              ],
            },
            last_name: {
              label: "Last Name",
              type: 'text',
              columns: {
                sm: {
                  container: 12,
                  label: 12,
                  wrapper: 12,
                },
                lg: {
                  container: 6,
                  label: 12,
                  wrapper: 12,
                },
              },
              fieldName: 'Last name',
              rules: [
                'required',
                'max:255',
              ],
            },
          },
          description: 'This must be your legal name.',
        },
        preferred_name_container: {
          label: "Preferred Name",
          type: 'group',
          schema: {
            preferred_name: {
              type: 'text',
              columns: {
                label: 12,
                wrapper: 12,
              },
              fieldName: 'preferredName',
              rules: [
                'max:255',
              ],
            },
          },
          description: '(Optional) Your preferred name.',
        },
        age: {
          name: "age",
          label: "Age",
          placeholder: 'Age',
          floating: false,
          type: 'text',
          inputType: 'number',
          columns: {
            lg: {
              container: 3,
              label: 12,
              wrapper: 12,
            },
          },
          rules: [
            'required',
            'max:100',
            'min:18',
            'numeric',
          ],
          fieldName: 'age',
        },
        phone: {
          label: "Phone Number",
          type: 'phone',
          rules: [
            'required',
          ],
          columns: {
            lg: {
              container: 3,
              label: 12,
              wrapper: 12,
            },
          },
          fieldName: 'Phone',
          placeholder: 'Phone',
          floating: false,
          allowIncomplete: true,
          unmask: true,
        },
        email: {
          label: "Email",
          placeholder: 'Email',
          floating: false,
          type: 'text',
          inputType: 'email',
          rules: [
            'required',
            'max:255',
            'email',
          ],
          columns: {
            lg: {
              container: 6,
              label: 12,
              wrapper: 12,
            },
          },
          fieldName: 'Email (School email preferred.',
        },
        country: {
          label: "Country of Residence",
          placeholder: 'Country of Residence',
          columns: {
            lg: {
              container: 6,
              label: 12,
              wrapper: 12,
            },
          },
          floating: false,
          type: 'select',
          search: true,
          native: false,
          inputType: 'search',
          autocomplete: 'enabled',
          items: countriesJson,
          rules: [
            'required',
          ],
        },
        uni: {
          label: "University",
          placeholder: 'University',
          columns: {
            lg: {
              container: 6,
              label: 12,
              wrapper: 12,
            },
          },
          floating: false,
          type: 'select',
          search: true,
          native: false,
          inputType: 'search',
          autocomplete: 'enabled',
          items: schoolsJson,
          rules: [
            'required',
          ],
          strict: false,
        },
        major: {
          label: 'Major',
          type: 'select',
          columns: {
            lg: {
              container: 6,
              label: 12,
              wrapper: 12,
            },
          },
          items: [
            {
              value: 'Business & Information Systems (BIS)',
              label: 'Business & Information Systems (BIS)',
            },
            {
              value: 'Computing & Business (CBUS)',
              label: 'Computing & Business (CBUS)',
            },
            {
              value: 'Computer Science (CS)',
              label: 'Computer Science (CS)',
            },
            {
              value: 'Computer Engineering (CoE)',
              label: 'Computer Engineering (CoE)',
            },
            {
              value: 'Data Science (DS)',
              label: 'Data Science (DS)',
            },
            {
              value: 'Human Computer Interaction (HCI)',
              label: 'Human Computer Interaction (HCI)',
            },
            {
              value: 'Information Systems (IS)',
              label: 'Information Systems (IS)',
            },
            {
              value: 'Information Technology (IT)',
              label: 'Information Technology (IT)',
            },
            {
              value: 'Software Engineering (SE)',
              label: 'Software Engineering (SE)',
            },
            {
              value: 'Web & Information Systems (WIS)',
              label: 'Web & Information Systems (WIS)',
            },
            {
              value: 'Mathematics or Statistics',
              label: 'Mathematics or Statistics',
            },
            {
              value: 'A natural science',
              label: 'A natural science (E.g. biology, chemistry, physics, etc.)',
            },
            {
              value: 'Another engineering discipline',
              label: 'Another engineering discipline (E.g. civil, electrical, mechanical, etc.)',
            },
            {
              value: 'Business Discipline',
              label: 'Business Discipline (E.g. accounting, finance, marketing, etc.)',
            },
            {
              value: 'Health Science',
              label: 'Health Science (E.g. nursing, pharmacy, radiology, etc.)',
            },
            {
              value: 'Fine arts or performing arts',
              label: 'Fine arts or performing arts (E.g. graphic design, music, studio art, etc.)',
            },
            {
              value: 'Other',
              label: 'Other',
            },
          ],
          view: "blocks",
          rules: [
            'required',
          ],
        },
        otherMajor: {
          type: "text",
          label: "Other Major",
          columns: {
            lg: {
              container: 6,
              label: 12,
              wrapper: 12,
            },
          },
          placeholder: "Other Major",
          floating: false,
          fieldName: 'Other Major',
          conditions: [['major', 'Other']],
          rules: [
            'required',
            'max:255',
          ],
        },

        container2: {
          type: "group",
          schema: {
            firsthack: {
              label: 'Is this your first Hackathon?',
              type: 'radiogroup',
              items: [
                {
                  value: 'Yes',
                  label: "Yes",
                },
                {
                  value: 'No',
                  label: 'No',
                },
              ],
              view: "tabs",
              columns: {
                lg: {
                  container: 6,
                  label: 12,
                  wrapper: 12,
                },
              },
              rules: [
                'required',
              ],
            },
            shirtsize: {
              label: 'Unisex T-Shirt Size',
              type: 'radiogroup',
              items: [
                {
                  value: 'XS',
                  label: 'XS',
                },
                {
                  value: 'S',
                  label: 'S',
                },
                {
                  value: 'M',
                  label: 'M',
                },
                {
                  value: 'L',
                  label: 'L',
                },
                {
                  value: 'XL',
                  label: 'XL',
                },
                {
                  value: '2XL',
                  label: '2XL',
                },
              ],
              view: "tabs",
              columns: {
                lg: {
                  container: 6,
                  label: 12,
                  wrapper: 12,
                },
              },
              rules: [
                'required',
              ],
            },
          },
        },
        container3: {
          type: "group",
          schema: {
            dietaryrestrictions: {
              label: 'Dietary Restrictions (Select all that apply)',
              type: 'checkboxgroup',
              items: [
                {
                  value: 'Vegetarian',
                  label: 'Vegetarian',
                },
                {
                  value: 'Vegan',
                  label: 'Vegan',
                },
                {
                  value: 'Kosher',
                  label: 'Kosher',
                },
                {
                  value: 'Halal',
                  label: 'Halal',
                },
                {
                  value: 'Gluten-free',
                  label: 'Gluten-free',
                },
              ],
              columns: {
                lg: {
                  container: 6,
                  label: 12,
                  wrapper: 12,
                },
              },
            },
            lvlofstudy: {
              label: 'Highest Level of Study',
              type: 'radiogroup',
              items: [
                {
                  value: 'Undergraduate University (2 year - Community college or similar)',
                  label: 'Undergraduate University (2 year - community college or similar)',
                },
                {
                  value: 'Undergraduate University (3+ year)',
                  label: 'Undergraduate University (3+ year)',
                },
                {
                  value: 'Graduate University (Masters, Professional, Doctoral, etc)',
                  label: 'Graduate University (Masters, Professional, Doctoral, etc)',
                },
                {
                  value: 'Code School / Bootcamp',
                  label: 'Code School / Bootcamp',
                },
                {
                  value: 'Other Vocational / Trade Program or Apprenticeship',
                  label: 'Other Vocational / Trade Program or Apprenticeship',
                },
                {
                  value: 'Prefer not to answer',
                  label: 'Prefer not to answer',
                },
                {
                  value: 'I’m not currently a student',
                  label: 'I’m not currently a student',
                },
              ],
              columns: {
                lg: {
                  container: 6,
                  label: 12,
                  wrapper: 12,
                },
              },
              rules: [
                'required',
              ],
            },
          },
        },
        h1: {
          type: 'static',
          tag: 'h1',
          content: 'MLH Partner Agreements',
        },
        divider_4: {
          type: 'static',
          tag: 'hr',
        },
        mlh_checkbox_0: {
          label: 'MLH Code of Conduct',
          type: 'checkbox',
          text: 'I have read and agree to the <a href=https://static.mlh.io/docs/mlh-code-of-conduct.pdf>MLH Code of Conduct</a>',
          rules: [
            'accepted',
          ],
        },
        mlh_checkbox_1: {
          label: 'MLH Terms and Conditions',
          type: 'checkbox',
          text: 'I authorize you to share my application/registration information with Major League Hacking for event administration, ranking, and MLH administration in-line with the <a href="https://mlh.io/privacy">MLH Privacy Policy</a>. I further agree to the terms of both the <a href="https://github.com/MLH/mlh-policies/blob/main/contest-terms.md">MLH Contest Terms and Conditions</a>  and the <a href="https://mlh.io/privacy">MLH Privacy Policy</a>.',
          rules: [
            'accepted',
          ],
        },
        mlh_checkbox_2: {
          label: 'MLH Email Policy',
          type: 'checkbox',
          text: '(Optional) I authorize MLH to send me occasional emails about relevant events, career opportunities, and community announcements.',
        },
        divider_3: {
          type: 'static',
          tag: 'hr',
        },
        divider_1: {
          type: 'static',
          tag: 'hr',
        },
        h1_2: {
          type: 'static',
          tag: 'h1',
          content: 'Optional Demographics',
        },
        p_1: {
          type: 'static',
          tag: 'p',
          content: '<div>This section is optional, but it helps us to get to know and serve our audience better!</div>',
        },
        minority: {
          type: 'radiogroup',
          label: 'Do you identify as part of an underrepresented group in the technology industry? ',
          items: [
            {
              value: 'Yes',
              label: 'Yes',
            },
            {
              value: 'No',
              label: 'No',
            },
            {
              value: 'Unsure',
              label: 'Unsure',
            },
            {
              value: 'Prefer Not to Answer',
              label: 'Prefer Not to Answer',
            },
          ],
        },
        gender: {
          type: 'radiogroup',
          label: 'Gender',
          items: [
            {
              value: 'Male',
              label: 'Male',
            },
            {
              value: 'Female',
              label: 'Female',
            },
            {
              value: 'Non-Binary',
              label: 'Non-Binary',
            },
            {
              value: 'Prefer to Self-Describe',
              label: 'Prefer to Self-Describe',
            },
            {
              value: 'Prefer Not to Answer',
              label: 'Prefer Not to Answer',
            },
          ],
        },
        race: {
          type: 'radiogroup',
          label: 'Race / Ethnicity ',
          items: [
            {
              value: 'Asian Indian',
              label: 'Asian Indian',
            },
            {
              value: 'Black or African',
              label: 'Black or African',
            },
            {
              value: 'Chinese',
              label: 'Chinese',
            },
            {
              value: 'Filipino',
              label: 'Filipino',
            },
            {
              value: 'Gaumanian or Chamorro',
              label: 'Gaumanian or Chamorro',
            },
            {
              value: 'Hispanic / Latino / Spanish Origin',
              label: 'Hispanic / Latino / Spanish Origin',
            },
            {
              value: 'Japenese',
              label: 'Japenese',
            },
            {
              value: 'Korean',
              label: 'Korean',
            },
            {
              value: 'Middle Eastern',
              label: 'Middle Eastern',
            },
            {
              value: 'Native Hawaiian',
              label: 'Native Hawaiian',
            },
            {
              value: 'Samoan',
              label: 'Samoan',
            },
            {
              value: 'Vietnamese',
              label: 'Vietnamese',
            },
            {
              value: 'White',
              label: 'White',
            },
            {
              value: 'Other Asian (Thai, Cambodian, etc)',
              label: 'Other Asian (Thai, Cambodian, etc)',
            },
            {
              value: 'Other Pacific Islander',
              label: 'Other Pacific Islander',
            },
            {
              value: 'Other',
              label: 'Other',
            },
            {
              value: 'Prefer Not to Answer',
              label: 'Prefer Not to Answer',
            },
          ],
        },
        divider_6: {
          type: 'static',
          tag: 'hr',
        },
        h1_1: {
          type: 'static',
          tag: 'h1',
          content: 'Networking (Optional)',
        },
        p: {
          type: 'static',
          tag: 'p',
          content: '<div>You can skip any questions in this section, but we would appreciate your response! We are working to help our sponsors and hackers connect. To that end, we will be sharing a resume booklet with our sponsors, as well as LinkedIn accounts. This information will be shared with all our sponsors.</div>',
        },
        divider_2: {
          type: 'static',
          tag: 'hr',
        },
        resume: {
          type: 'file',
          label: 'Resume (Attach your resume as a PDF',
          accept: 'pdf,application/pdf',
          rules: [
            'max:10000',
          ],
          upload: false,
          uploadTempEndpoint: false,
          auto: false,
          uploadUrl: null,
        },
        linkedin: {
          type: 'text',
          label: 'LinkedIn Account URL',
          inputType: 'url',
        },
        divider_5: {
          type: 'static',
          tag: 'hr',
        },
      },
    }
  }),
  methods: {
    handleResponse(response, form$) {
      this.response = response;
    },
    handleError(error, details, form$) {
      this.response = error;
    }
  }
}
</script>

<style>
.vf-registration *,
.vf-registration *:before,
.vf-registration *:after,
.vf-registration:root {
  --vf-primary: #058168;
  --vf-primary-darker: #035f4d;
  --vf-color-on-primary: #ffffff;
  --vf-danger: #c03737;
  --vf-danger-lighter: #fee2e2;
  --vf-success: #058168;
  --vf-success-lighter: #d1fae5;
  --vf-gray-50: #f9fafb;
  --vf-gray-100: #f3f4f6;
  --vf-gray-200: #e5e7eb;
  --vf-gray-300: #d1d5db;
  --vf-gray-400: #9ca3af;
  --vf-gray-500: #6b7280;
  --vf-gray-600: #4b5563;
  --vf-gray-700: #374151;
  --vf-gray-800: #1f2937;
  --vf-gray-900: #111827;
  --vf-dark-50: #EFEFEF;
  --vf-dark-100: #DCDCDC;
  --vf-dark-200: #BDBDBD;
  --vf-dark-300: #A0A0A0;
  --vf-dark-400: #848484;
  --vf-dark-500: #737373;
  --vf-dark-600: #393939;
  --vf-dark-700: #323232;
  --vf-dark-800: #262626;
  --vf-dark-900: #191919;
  --vf-ring-width: 2px;
  --vf-ring-color: #07bf9b66;
  --vf-link-color: var(--vf-color-success);
  --vf-link-decoration: underline;
  --vf-font-size: 1rem;
  --vf-font-size-sm: 0.875rem;
  --vf-font-size-lg: 1rem;
  --vf-font-size-small: 0.875rem;
  --vf-font-size-small-sm: 0.8125rem;
  --vf-font-size-small-lg: 0.875rem;
  --vf-font-size-h1: 2.125rem;
  --vf-font-size-h1-sm: 2.125rem;
  --vf-font-size-h1-lg: 2.125rem;
  --vf-font-size-h2: 1.875rem;
  --vf-font-size-h2-sm: 1.875rem;
  --vf-font-size-h2-lg: 1.875rem;
  --vf-font-size-h3: 1.5rem;
  --vf-font-size-h3-sm: 1.5rem;
  --vf-font-size-h3-lg: 1.5rem;
  --vf-font-size-h4: 1.25rem;
  --vf-font-size-h4-sm: 1.25rem;
  --vf-font-size-h4-lg: 1.25rem;
  --vf-font-size-h1-mobile: 1.5rem;
  --vf-font-size-h1-mobile-sm: 1.5rem;
  --vf-font-size-h1-mobile-lg: 1.5rem;
  --vf-font-size-h2-mobile: 1.25rem;
  --vf-font-size-h2-mobile-sm: 1.25rem;
  --vf-font-size-h2-mobile-lg: 1.25rem;
  --vf-font-size-h3-mobile: 1.125rem;
  --vf-font-size-h3-mobile-sm: 1.125rem;
  --vf-font-size-h3-mobile-lg: 1.125rem;
  --vf-font-size-h4-mobile: 1rem;
  --vf-font-size-h4-mobile-sm: 1rem;
  --vf-font-size-h4-mobile-lg: 1rem;
  --vf-font-size-blockquote: 1rem;
  --vf-font-size-blockquote-sm: 0.875rem;
  --vf-font-size-blockquote-lg: 1rem;
  --vf-line-height: 1.5rem;
  --vf-line-height-sm: 1.25rem;
  --vf-line-height-lg: 1.5rem;
  --vf-line-height-small: 1.25rem;
  --vf-line-height-small-sm: 1.125rem;
  --vf-line-height-small-lg: 1.25rem;
  --vf-line-height-headings: 1.2;
  --vf-line-height-headings-sm: 1.2;
  --vf-line-height-headings-lg: 1.2;
  --vf-line-height-blockquote: 1.5rem;
  --vf-line-height-blockquote-sm: 1.25rem;
  --vf-line-height-blockquote-lg: 1.5rem;
  --vf-letter-spacing: 0px;
  --vf-letter-spacing-sm: 0px;
  --vf-letter-spacing-lg: 0px;
  --vf-letter-spacing-small: 0px;
  --vf-letter-spacing-small-sm: 0px;
  --vf-letter-spacing-small-lg: 0px;
  --vf-letter-spacing-headings: 0px;
  --vf-letter-spacing-headings-sm: 0px;
  --vf-letter-spacing-headings-lg: 0px;
  --vf-letter-spacing-blockquote: 0px;
  --vf-letter-spacing-blockquote-sm: 0px;
  --vf-letter-spacing-blockquote-lg: 0px;
  --vf-gutter: 1rem;
  --vf-gutter-sm: 0.5rem;
  --vf-gutter-lg: 1rem;
  --vf-min-height-input: 2.375rem;
  --vf-min-height-input-sm: 2.125rem;
  --vf-min-height-input-lg: 2.875rem;
  --vf-py-input: 0.375rem;
  --vf-py-input-sm: 0.375rem;
  --vf-py-input-lg: 0.625rem;
  --vf-px-input: 0.75rem;
  --vf-px-input-sm: 0.5rem;
  --vf-px-input-lg: 0.875rem;
  --vf-py-btn: 0.375rem;
  --vf-py-btn-sm: 0.375rem;
  --vf-py-btn-lg: 0.625rem;
  --vf-px-btn: 0.875rem;
  --vf-px-btn-sm: 0.75rem;
  --vf-px-btn-lg: 1.25rem;
  --vf-py-btn-small: 0.25rem;
  --vf-py-btn-small-sm: 0.25rem;
  --vf-py-btn-small-lg: 0.375rem;
  --vf-px-btn-small: 0.625rem;
  --vf-px-btn-small-sm: 0.625rem;
  --vf-px-btn-small-lg: 0.75rem;
  --vf-py-group-tabs: 0.375rem;
  --vf-py-group-tabs-sm: 0.375rem;
  --vf-py-group-tabs-lg: 0.625rem;
  --vf-px-group-tabs: 0.75rem;
  --vf-px-group-tabs-sm: 0.5rem;
  --vf-px-group-tabs-lg: 0.875rem;
  --vf-py-group-blocks: 0.75rem;
  --vf-py-group-blocks-sm: 0.625rem;
  --vf-py-group-blocks-lg: 0.875rem;
  --vf-px-group-blocks: 1rem;
  --vf-px-group-blocks-sm: 1rem;
  --vf-px-group-blocks-lg: 1rem;
  --vf-py-tag: 0px;
  --vf-py-tag-sm: 0px;
  --vf-py-tag-lg: 0px;
  --vf-px-tag: 0.4375rem;
  --vf-px-tag-sm: 0.4375rem;
  --vf-px-tag-lg: 0.4375rem;
  --vf-py-slider-tooltip: 0.125rem;
  --vf-py-slider-tooltip-sm: 0.0625rem;
  --vf-py-slider-tooltip-lg: 0.1875rem;
  --vf-px-slider-tooltip: 0.375rem;
  --vf-px-slider-tooltip-sm: 0.3125rem;
  --vf-px-slider-tooltip-lg: 0.5rem;
  --vf-py-blockquote: 0.25rem;
  --vf-py-blockquote-sm: 0.25rem;
  --vf-py-blockquote-lg: 0.25rem;
  --vf-px-blockquote: 0.75rem;
  --vf-px-blockquote-sm: 0.75rem;
  --vf-px-blockquote-lg: 0.75rem;
  --vf-py-hr: 0.25rem;
  --vf-space-addon: 0px;
  --vf-space-addon-sm: 0px;
  --vf-space-addon-lg: 0px;
  --vf-space-checkbox: 0.375rem;
  --vf-space-checkbox-sm: 0.375rem;
  --vf-space-checkbox-lg: 0.375rem;
  --vf-space-tags: 0.1875rem;
  --vf-space-tags-sm: 0.1875rem;
  --vf-space-tags-lg: 0.1875rem;
  --vf-space-static-tag-1: 1rem;
  --vf-space-static-tag-2: 2rem;
  --vf-space-static-tag-3: 3rem;
  --vf-floating-top: 0rem;
  --vf-floating-top-sm: 0rem;
  --vf-floating-top-lg: 0.6875rem;
  --vf-bg-input: #ffffff;
  --vf-bg-input-hover: #ffffff;
  --vf-bg-input-focus: #ffffff;
  --vf-bg-input-danger: #ffffff;
  --vf-bg-input-success: #ffffff;
  --vf-bg-checkbox: #ffffff;
  --vf-bg-checkbox-hover: #ffffff;
  --vf-bg-checkbox-focus: #ffffff;
  --vf-bg-checkbox-danger: #ffffff;
  --vf-bg-checkbox-success: #ffffff;
  --vf-bg-disabled: var(--vf-gray-200);
  --vf-bg-selected: #1118270d;
  --vf-bg-passive: var(--vf-gray-300);
  --vf-bg-icon: var(--vf-gray-500);
  --vf-bg-danger: var(--vf-danger-lighter);
  --vf-bg-success: var(--vf-success-lighter);
  --vf-bg-tag: var(--vf-primary);
  --vf-bg-slider-handle: var(--vf-primary);
  --vf-bg-toggle-handle: #ffffff;
  --vf-bg-date-head: var(--vf-gray-100);
  --vf-bg-addon: #ffffff00;
  --vf-bg-btn: var(--vf-primary);
  --vf-bg-btn-danger: var(--vf-danger);
  --vf-bg-btn-secondary: var(--vf-gray-200);
  --vf-bg-table-header: var(--vf-gray-100);
  --vf-color-input: var(--vf-gray-800);
  --vf-color-input-hover: var(--vf-gray-800);
  --vf-color-input-focus: var(--vf-gray-800);
  --vf-color-input-danger: var(--vf-gray-800);
  --vf-color-input-success: var(--vf-gray-800);
  --vf-color-disabled: var(--vf-gray-400);
  --vf-color-passive: var(--vf-gray-700);
  --vf-color-muted: var(--vf-gray-500);
  --vf-color-floating: var(--vf-gray-500);
  --vf-color-floating-focus: var(--vf-gray-500);
  --vf-color-floating-success: var(--vf-gray-500);
  --vf-color-floating-danger: var(--vf-gray-500);
  --vf-color-danger: var(--vf-danger);
  --vf-color-success: var(--vf-success);
  --vf-color-tag: var(--vf-color-on-primary);
  --vf-color-addon: var(--vf-gray-800);
  --vf-color-date-head: var(--vf-gray-700);
  --vf-color-btn: var(--vf-color-on-primary);
  --vf-color-btn-danger: #ffffff;
  --vf-color-btn-secondary: var(--vf-gray-700);
  --vf-color-table-header: var(--vf-gray-800);
  --vf-border-color-input: var(--vf-gray-300);
  --vf-border-color-input-hover: var(--vf-gray-300);
  --vf-border-color-input-focus: var(--vf-primary);
  --vf-border-color-input-danger: var(--vf-gray-300);
  --vf-border-color-input-success: var(--vf-gray-300);
  --vf-border-color-checkbox: var(--vf-gray-300);
  --vf-border-color-checkbox-focus: var(--vf-primary);
  --vf-border-color-checkbox-hover: var(--vf-gray-300);
  --vf-border-color-checkbox-danger: var(--vf-gray-300);
  --vf-border-color-checkbox-success: var(--vf-gray-300);
  --vf-border-color-checked: var(--vf-primary);
  --vf-border-color-passive: var(--vf-gray-300);
  --vf-border-color-slider-tooltip: var(--vf-primary);
  --vf-border-color-tag: var(--vf-primary);
  --vf-border-color-btn: var(--vf-primary);
  --vf-border-color-btn-danger: var(--vf-danger);
  --vf-border-color-btn-secondary: var(--vf-gray-200);
  --vf-border-color-blockquote: var(--vf-gray-300);
  --vf-border-color-hr: var(--vf-gray-300);
  --vf-border-color-signature-hr: var(--vf-gray-300);
  --vf-border-color-table: var(--vf-gray-300);
  --vf-border-width-input-t: 1px;
  --vf-border-width-input-r: 1px;
  --vf-border-width-input-b: 1px;
  --vf-border-width-input-l: 1px;
  --vf-border-width-radio-t: 1px;
  --vf-border-width-radio-r: 1px;
  --vf-border-width-radio-b: 1px;
  --vf-border-width-radio-l: 1px;
  --vf-border-width-checkbox-t: 1px;
  --vf-border-width-checkbox-r: 1px;
  --vf-border-width-checkbox-b: 1px;
  --vf-border-width-checkbox-l: 1px;
  --vf-border-width-dropdown: 1px;
  --vf-border-width-btn: 1px;
  --vf-border-width-toggle: 0.125rem;
  --vf-border-width-tag: 1px;
  --vf-border-width-blockquote: 3px;
  --vf-border-width-table: 1px;
  --vf-shadow-input: 0px 0px 0px 0px rgba(0, 0, 0, 0);
  --vf-shadow-input-hover: 0px 0px 0px 0px rgba(0, 0, 0, 0);
  --vf-shadow-input-focus: 0px 0px 0px 0px rgba(0, 0, 0, 0);
  --vf-shadow-handles: 0px 0px 0px 0px rgba(0, 0, 0, 0);
  --vf-shadow-handles-hover: 0px 0px 0px 0px rgba(0, 0, 0, 0);
  --vf-shadow-handles-focus: 0px 0px 0px 0px rgba(0, 0, 0, 0);
  --vf-shadow-btn: 0px 0px 0px 0px rgba(0, 0, 0, 0);
  --vf-shadow-dropdown: 0px 0px 0px 0px rgba(0, 0, 0, 0);
  --vf-radius-input: 0.25rem;
  --vf-radius-input-sm: 0.25rem;
  --vf-radius-input-lg: 0.25rem;
  --vf-radius-btn: 0.25rem;
  --vf-radius-btn-sm: 0.25rem;
  --vf-radius-btn-lg: 0.25rem;
  --vf-radius-small: 0.25rem;
  --vf-radius-small-sm: 0.25rem;
  --vf-radius-small-lg: 0.25rem;
  --vf-radius-large: 0.25rem;
  --vf-radius-large-sm: 0.25rem;
  --vf-radius-large-lg: 0.25rem;
  --vf-radius-tag: 0.25rem;
  --vf-radius-tag-sm: 0.25rem;
  --vf-radius-tag-lg: 0.25rem;
  --vf-radius-checkbox: 0.25rem;
  --vf-radius-checkbox-sm: 0.25rem;
  --vf-radius-checkbox-lg: 0.25rem;
  --vf-radius-slider: 0.25rem;
  --vf-radius-slider-sm: 0.25rem;
  --vf-radius-slider-lg: 0.25rem;
  --vf-radius-image: 0.25rem;
  --vf-radius-image-sm: 0.25rem;
  --vf-radius-image-lg: 0.25rem;
  --vf-radius-gallery: 0.25rem;
  --vf-radius-gallery-sm: 0.25rem;
  --vf-radius-gallery-lg: 0.25rem;
  --vf-checkbox-size: 1rem;
  --vf-checkbox-size-sm: 0.875rem;
  --vf-checkbox-size-lg: 1rem;
  --vf-gallery-size: 6rem;
  --vf-gallery-size-sm: 5rem;
  --vf-gallery-size-lg: 7rem;
  --vf-toggle-width: 3rem;
  --vf-toggle-width-sm: 2.75rem;
  --vf-toggle-width-lg: 3rem;
  --vf-toggle-height: 1.25rem;
  --vf-toggle-height-sm: 1rem;
  --vf-toggle-height-lg: 1.25rem;
  --vf-slider-height: 0.375rem;
  --vf-slider-height-sm: 0.3125rem;
  --vf-slider-height-lg: 0.5rem;
  --vf-slider-height-vertical: 20rem;
  --vf-slider-height-vertical-sm: 20rem;
  --vf-slider-height-vertical-lg: 20rem;
  --vf-slider-handle-size: 1rem;
  --vf-slider-handle-size-sm: 0.875rem;
  --vf-slider-handle-size-lg: 1.25rem;
  --vf-slider-tooltip-distance: 0.5rem;
  --vf-slider-tooltip-distance-sm: 0.375rem;
  --vf-slider-tooltip-distance-lg: 0.5rem;
  --vf-slider-tooltip-arrow-size: 0.3125rem;
  --vf-slider-tooltip-arrow-size-sm: 0.3125rem;
  --vf-slider-tooltip-arrow-size-lg: 0.3125rem;
}

.vf-description {
  text-align: left;
}

.vf-file-preview-upload {
  display: none;
}

*[aria-labelledby="container__label"] {
  background-color: #eef3f7;
  padding: 8px;
  border-radius: 8px;
}

.vf-registration *,
.form-submitted * {
  font-family: sans-serif;
}

.form-submitted h1 {
  font-size: var(--vf-font-size-h1);
}

.form-submitted p {
  font-size: var(--vf-font-size);
  margin-bottom: 24px;
}

.vf-element-layout-outer-wrapper {
  margin: 4px;
}
</style>
